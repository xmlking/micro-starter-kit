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
* - MICRO_PROXY_ADDRESS="localhost:8888"
* You can also run this test againest your local running service. i.e., `go run ./srv/greeter`
**/
func init() {
	// if start proxy and testing with MICRO_PROXY_ADDRESS="localhost:8888"
	greeter = greeterPB.NewGreeterService("greetersrv", grpc.NewClient())
	// if start GreeterService with `make run-greeter ARGS="--server_address=localhost:8080"`
	// greeter = greeterPB.NewGreeterService("localhost", grpc.NewClient(client.Selector(static.NewSelector())))
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
