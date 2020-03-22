// e2e, black-box testing
package e2e

import (
	"context"
	"testing"

	"github.com/micro/go-micro/v2/client"
	"github.com/stretchr/testify/assert"

	// "github.com/xmlking/micro-starter-kit/shared/micro/client/selector/static"
	greeterPB "github.com/xmlking/micro-starter-kit/service/greeter/proto/greeter"
	"github.com/xmlking/micro-starter-kit/shared/constants"
)

var (
	greeterSrvClient greeterPB.GreeterService
)

/**
* set envelopment variables for CI e2e tests with `memory` registry.
* - export MICRO_REGISTRY=memory
* - export MICRO_SELECTOR=static
* (Or) Set envelopment variables for CI e2e tests via gRPC Proxy
* - MICRO_PROXY_ADDRESS="localhost:8081"
* You can also run this test againest your local running service with mDNS. i.e., `make run-greeter`
**/
func init() {
	greeterSrvClient = greeterPB.NewGreeterService(constants.GREETER_SERVICE, client.DefaultClient)
}

func TestGreeter_Hello_E2E(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test")
	}

	rsp, err := greeterSrvClient.Hello(context.TODO(), &greeterPB.HelloRequest{Name: "Sumo"})
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, rsp.Msg, "Hello Sumo")
}
