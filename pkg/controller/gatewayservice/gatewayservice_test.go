package gatewayservice

import (
	"github.com/devcsrj/gravitee-k8operator/pkg/apis/devcsrj/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"net/http"
	"os"
	"path/filepath"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"
)

const (
	apimUrl   = "http://apim.example.com"
	apimToken = "YXBpbXVzZXI6YXBpbXB3Cg=="
)

var _ = Describe("Gatewayservice", func() {
	var svc *corev1.Service
	var gatewaySvc *v1alpha1.GatewayService
	var reconciler *ReconcileGatewayService
	var namespace = "gravitee"

	BeforeEach(func() {
		gock.DisableNetworking()

		svc = &corev1.Service{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Service",
				APIVersion: "core/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "petstore",
				Namespace: namespace,
				Labels: map[string]string{
					"app": "petstore",
				},
			},
			Spec: corev1.ServiceSpec{
				Ports: []corev1.ServicePort{
					{
						Name: "http",
						Port: 80,
						TargetPort: intstr.IntOrString{
							IntVal: 8080,
						},
					},
				},
			},
			Status: corev1.ServiceStatus{},
		}

		gatewaySvc = &v1alpha1.GatewayService{
			TypeMeta: metav1.TypeMeta{
				Kind:       "GatewayService",
				APIVersion: "devcsrj/v1alpha1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "demo-gatewayservice",
				Namespace: namespace,
			},
			Spec: v1alpha1.GatewayServiceSpec{
				Selector: map[string]string{
					"app": "petstore",
				},
				OasPath: "/openapi",
			},
			Status: v1alpha1.GatewayServiceStatus{},
		}
	})

	Describe("Reconciling", func() {
		BeforeEach(func() {
			scheme := runtime.NewScheme()
			scheme.AddKnownTypes(v1alpha1.SchemeGroupVersion, gatewaySvc)
			Expect(corev1.AddToScheme(scheme)).To(Succeed())

			client := fake.NewFakeClientWithScheme(scheme, svc, gatewaySvc)

			httpClient := http.Client{
				Timeout: 1 * time.Second,
			}

			reconciler = &ReconcileGatewayService{
				client:     client,
				scheme:     scheme,
				httpClient: httpClient,
				apimToken:  apimToken,
				apimUrl:    apimUrl,
			}
		})

		Context("a freshly deployed Service", func() {
			It("should publish API", func() {
				defer gock.Off()

				importResp, _ := os.Open(filepath.Join("testdata", "import-swagger.200-imported.json"))
				defer importResp.Close()
				gock.New(apimUrl).
					Post("/management/apis/import/swagger").
					MatchHeader("Authorization", "Basic "+apimToken).
					MatchHeader("Accept", "application/json").
					MatchType("json").
					//JSON(map[string]string{"type": "URL", "payload": "http://petstore/openapi"}).
					Reply(200).
					Type("json").
					Body(importResp)

				req := reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: namespace,
						Name:      "demo-gatewayservice",
					},
				}
				result, err := reconciler.Reconcile(req)

				Expect(result).ToNot(BeNil())
				Expect(err).To(BeNil())
				Expect(gock.IsDone()).To(BeTrue())
			})
		})

		Context("invalid credentials", func() {
			It("should return an error", func() {
				defer gock.Off()

				importResp, _ := os.Open(filepath.Join("testdata", "import-swagger.401-unauthorized.txt"))
				defer importResp.Close()
				gock.New(apimUrl).
					Post("/management/apis/import/swagger").
					Reply(401).
					Body(importResp)

				req := reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: namespace,
						Name:      "demo-gatewayservice",
					},
				}
				result, err := reconciler.Reconcile(req)

				Expect(result).ToNot(BeNil())
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("petstore - credentials was rejected by gravitee"))
				Expect(gock.IsDone()).To(BeTrue())
			})
		})

		Context("bad swagger spec", func() {
			It("should return an error", func() {
				defer gock.Off()

				importResp, _ := os.Open(filepath.Join("testdata", "import-swagger.400-bad-format.json"))
				defer importResp.Close()
				gock.New(apimUrl).
					Post("/management/apis/import/swagger").
					Reply(400).
					Body(importResp)

				req := reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: namespace,
						Name:      "demo-gatewayservice",
					},
				}
				result, err := reconciler.Reconcile(req)

				Expect(result).ToNot(BeNil())
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("petstore - Swagger descriptor: bad format"))
				Expect(gock.IsDone()).To(BeTrue())
			})
		})

		Context("path already covered", func() {
			It("should return an error", func() {
				defer gock.Off()

				importResp, _ := os.Open(filepath.Join("testdata", "import-swagger.400-already-covered.json"))
				defer importResp.Close()
				gock.New(apimUrl).
					Post("/management/apis/import/swagger").
					Reply(400).
					Body(importResp)

				req := reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: namespace,
						Name:      "demo-gatewayservice",
					},
				}
				result, err := reconciler.Reconcile(req)

				Expect(result).ToNot(BeNil())
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("petstore - The path [/v3/] is already covered by an other API."))
				Expect(gock.IsDone()).To(BeTrue())
			})
		})
	})
})
