package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	account "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
)

type Account struct{}

func (e *Account) Handle(ctx context.Context, msg *account.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *account.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
