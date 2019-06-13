package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	account "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
)

// AccountSubscriber struct
type AccountSubscriber struct{}

// Handle example
func (e *AccountSubscriber) Handle(ctx context.Context, user *account.User) error {
	log.Logf("Handler Received message: %v", user.Username)
	return nil
}

// Handler function
func Handler(ctx context.Context, user *account.User) error {
	log.Logf("Function Received message: %v", user.Username)
	return nil
}
