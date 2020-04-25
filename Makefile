ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

GRAVITEE_VERSION ?= 1.22

gravitee-sdk:
	openapi-generator generate -i https://raw.githubusercontent.com/gravitee-io/gravitee-docs/master/apim/api/$(GRAVITEE_VERSION)/swagger.json\
		--generator-name go\
		--additional-properties "packageName=graviteeapim"\
		--config openapi-generator-opts.json\
		--git-user-id devcsrj\
		--git-repo-id gravitee-operator\
		--skip-validate-spec\
		--output pkg/gravitee-apim-sdk
