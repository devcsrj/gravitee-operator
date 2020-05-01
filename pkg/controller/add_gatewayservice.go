package controller

import (
	"github.com/devcsrj/gravitee-operator/pkg/controller/gatewayservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, gatewayservice.Add)
}
