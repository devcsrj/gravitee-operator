ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

generate:
	operator-sdk generate crds
	operator-sdk generate k8s
image:
	operator-sdk build devcsrj/gravitee-operator