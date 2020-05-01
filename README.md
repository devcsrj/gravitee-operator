# Gravitee | Operator

![Travis (.org)](https://img.shields.io/travis/devcsrj/gravitee-k8operator)

A [kubernetes operator](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) responsible 
for publishing [OpenAPI specifications](https://swagger.io/specification/) hosted by a [Service](https://kubernetes.io/docs/concepts/services-networking/service/) 
to [Gravitee's API Management](https://gravitee.io/products/apim/) platform.

## Example

```yaml
apiVersion: devcsrj.gravitee.io/v1alpha1
kind: GatewayService
metadata:
  name: gravitee
  namespace: default
spec:
  selector:
    app: petstore
  oasPath: /openapi
```

When applied to `kubectl`, this will lookup all `kind: Service`s in the `default` namespace, and 
expect that all `Pod`s behind this service will return a valid OpenAPI specification under 
the path `/openapi`.

## Installation

TODO

