// e2e, black-box testing
package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	emailerPB "github.com/xmlking/micro-starter-kit/srv/emailer/proto/emailer"
)

var (
	publisher micro.Publisher
	topic     = "emailersrv"
)

func init() {
	publisher = micro.NewPublisher(topic, client.DefaultClient)
}

func TestEmailSubscriber_Handle_E2E(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test")
	}

	msg := &emailerPB.Message{To: "sumo@demo.com"}

	err := publisher.Publish(context.TODO(), msg)
	if err != nil {
		t.Error(err)
	}
	time.Sleep(2 * time.Second)
	t.Logf("Successfully published to: %s", topic)
}
