package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	echo "github.com/xmlking/micro-starter-kit/srv/echo/proto/echo"
)

type Echo struct{}

func (e *Echo) Handle(ctx context.Context, msg *echo.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *echo.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
