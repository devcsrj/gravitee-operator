package gatewayservice

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	errors2 "errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	devcsrjv1alpha1 "github.com/devcsrj/gravitee-operator/pkg/apis/devcsrj/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const (
	apimMetaApiId = "apim.gravitee.io/apiId"
)

var log = logf.Log.WithName("controller_gatewayservice")

// Add creates a new GatewayService Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	apimUrl := os.Getenv("GRAVITEE_APIM_URL")
	if len(apimUrl) <= 0 {
		panic("missing 'GRAVITEE_APIM_URL' in the environment")
	}
	apimUsername := os.Getenv("GRAVITEE_APIM_USERNAME")
	apimPassword := os.Getenv("GRAVITEE_APIM_PASSWORD")
	token := base64.StdEncoding.EncodeToString([]byte(apimUsername + ":" + apimPassword))
	apimTimeout := 5 * time.Second
	if val, ok := os.LookupEnv("GRAVITEE_APIM_TIMEOUTMS"); ok {
		parsed, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			apimTimeout = time.Duration(parsed) * time.Millisecond
		}
	}

	httpClient := http.Client{
		Timeout: apimTimeout,
	}
	return &ReconcileGatewayService{
		client:     mgr.GetClient(),
		scheme:     mgr.GetScheme(),
		httpClient: httpClient,
		apimToken:  token,
		apimUrl:    apimUrl,
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("gatewayservice-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource GatewayService
	err = c.Watch(&source.Kind{Type: &devcsrjv1alpha1.GatewayService{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner GatewayService
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &devcsrjv1alpha1.GatewayService{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileGatewayService implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileGatewayService{}

// ReconcileGatewayService reconciles a GatewayService object
type ReconcileGatewayService struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client     client.Client
	scheme     *runtime.Scheme
	httpClient http.Client
	apimToken  string
	apimUrl    string
}

// Reconcile reads that state of the cluster for a GatewayService object and makes changes based on the state read
// and what is in the GatewayService.Spec. It fetches all matching `Service`s, and retrieves their
// http-exposed OpenAPI specification. The OAS is then forwarded to Gravitee.
//
// As of the moment, this only supports `Service`s of type ClusterIP
func (r *ReconcileGatewayService) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling GatewayService")

	// Fetch the GatewayService instance
	instance := &devcsrjv1alpha1.GatewayService{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	ctx := context.TODO()
	services, err := r.matchingServices(ctx, request, instance)
	if err != nil || len(services.Items) <= 0 {
		return reconcile.Result{}, err
	}

	// Publish each matched Service into the Gravitee
	errs := make(map[string]error, 0)
	for _, item := range services.Items {
		reqLogger.Info("Publishing Service to Gravitee",
			"Service.Namespace", item.Namespace,
			"Service.Name", item.Name,
			"GatewayService.Name", instance.Name)

		err := r.ensureApiIsPublished(ctx, item, instance)
		if err != nil {
			errs[item.Name] = err
		}
	}

	if len(errs) > 0 {
		var b strings.Builder
		b.WriteString("failed to publish the specifications from the following:")
		b.WriteString("\n")
		for k, v := range errs {
			b.WriteString(fmt.Sprintf("\t* %s - %s", k, v.Error()))
		}
		return reconcile.Result{}, errors2.New(b.String())
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileGatewayService) matchingServices(context context.Context, request reconcile.Request, gatewaySvc *devcsrjv1alpha1.GatewayService) (*corev1.ServiceList, error) {
	services := corev1.ServiceList{
		Items: make([]corev1.Service, 0),
	}
	opts := []client.ListOption{
		client.InNamespace(request.Namespace),
		client.MatchingLabels(gatewaySvc.Spec.Selector),
	}
	err := r.client.List(context, &services, opts...)
	return &services, err
}

func (r *ReconcileGatewayService) ensureApiIsPublished(ctx context.Context, item corev1.Service, gatewaySvc *devcsrjv1alpha1.GatewayService) error {
	existingId := item.Annotations[apimMetaApiId]
	if len(existingId) > 0 {
		// TODO has existing api id
	}

	// New API! Let's publish
	_, err := r.doPublish(ctx, item, gatewaySvc)
	if err != nil {
		return err
	}

	// TODO add id to service annotations

	return nil
}

func (r *ReconcileGatewayService) doPublish(ctx context.Context, item corev1.Service, gatewaySvc *devcsrjv1alpha1.GatewayService) (*apimApiId, error) {
	for _, port := range item.Spec.Ports {
		// Trim leading slash
		segment := gatewaySvc.Spec.OasPath
		for len(segment) > 0 && segment[0] == '/' {
			segment = segment[1:]
		}

		host := "http//" + item.Name
		if port.Port != 80 {
			host = fmt.Sprintf("%s:%d", host, port.Port)
		}
		reqPayload, _ := json.Marshal(map[string]string{
			"type":    "URL",
			"payload": fmt.Sprintf("%s/%s", host, segment),
		})
		parsedUrl, _ := url.Parse(r.apimUrl + "/management/apis/import/swagger")
		request := http.Request{
			Method: "POST",
			URL:    parsedUrl,
			Header: map[string][]string{
				"Authorization": {"Basic " + r.apimToken},
				"Content-Type":  {"application/json"},
				"Accept":        {"application/json"},
			},
			Body: ioutil.NopCloser(bytes.NewReader(reqPayload)),
		}

		response, err := r.httpClient.Do(&request)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()
		if response.StatusCode == 401 {
			return nil, errors2.New("credentials was rejected by gravitee")
		}

		// We deserialize and take the 'id' field
		var payload map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&payload)
		if err != nil {
			return nil, err
		}

		if response.StatusCode/100 == 4 {
			return nil, errors2.New(payload["message"].(string))
		}

		id := payload["id"].(string)
		return &apimApiId{id}, nil
	}

	return nil, errors2.New("no reachable port exposed")
}

type apimApiId struct {
	value string
}
