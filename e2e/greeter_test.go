// e2e, black-box testing
package e2e

import (
	"context"
	"testing"

	// "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/grpc"
	"github.com/stretchr/testify/assert"

	// "github.com/xmlking/micro-starter-kit/shared/micro/client/selector/static"
	greeterPB "github.com/xmlking/micro-starter-kit/srv/greeter/proto/greeter"
)

var (
	greeter greeterPB.GreeterService
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
	greeter = greeterPB.NewGreeterService("greetersrv", grpc.NewClient())
}

func TestGreeter_Hello_E2E(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test")
	}

	rsp, err := greeter.Hello(context.TODO(), &greeterPB.Request{Name: "Sumo"})
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, rsp.Msg, "Hello Sumo")
}
