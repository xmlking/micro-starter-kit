package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	emailerPB "github.com/xmlking/micro-starter-kit/srv/emailer/proto/emailer"
)

// Emailer is Subscriber
type Emailer struct{}

// Handle is a method to send emails
func (e *Emailer) Handle(ctx context.Context, msg *emailerPB.Message) error {
	log.Logf("Sending email to: %s from: %s subject: %s body: %s", msg.To, msg.From, msg.Subject, msg.Body)
	return nil
}

// Handler is a function to send emails
func Handler(ctx context.Context, msg *emailerPB.Message) error {
	log.Logf("Sending email to: %s from: %s subject: %s body: %s", msg.To, msg.From, msg.Subject, msg.Body)
	return nil
}
