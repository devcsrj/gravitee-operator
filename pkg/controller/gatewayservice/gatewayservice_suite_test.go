package gatewayservice

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGatewayservice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gatewayservice Suite")
}
