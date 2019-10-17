// e2e, black-box testing
package greeter_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/micro/go-micro/client/grpc"
	proto "github.com/xmlking/micro-starter-kit/srv/greeter/proto/greeter"
)

var (
	greeter proto.GreeterService
)

/**
 * set envelopment variables for CI e2e tests
 * - MICRO_REGISTRY=consul
 * - MICRO_REGISTRY_ADDRESS="$(CONSUL_SRV_ENDPOINT):8500"
 * - MICRO_REGISTER_TTL="10"
 * - MICRO_REGISTER_INTERVAL="5"
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
