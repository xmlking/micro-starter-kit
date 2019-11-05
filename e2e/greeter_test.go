// e2e, black-box testing
package e2e

import (
	"context"
	"testing"

	"github.com/micro/go-micro/client/grpc"
	"github.com/stretchr/testify/assert"
	proto "github.com/xmlking/micro-starter-kit/srv/greeter/proto/greeter"
)

var (
	greeter proto.GreeterService
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
	greeter = proto.NewGreeterService("greetersrv", grpc.NewClient())
}

func TestGreeter_Hello_E2E(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test")
	}

	rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name: "Sumo"})
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, rsp.Msg, "Hello Sumo")
}
