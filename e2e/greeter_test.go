// e2e, black-box testing
package greeter_test

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
* set envelopment variables for CI e2e tests with memory registry
* - MICRO_REGISTRY=memory
* - MICRO_SELECTOR=static
* - MICRO_PROXY_ADDRESS="localhost:8081"
* set envelopment variables for CI e2e tests with etcd registry
* - MICRO_REGISTRY=etcd
* - MICRO_REGISTRY_ADDRESS="prod-etcd-cluster-v1-client"
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
