package subscriber

import (
	"context"

	log "github.com/sirupsen/logrus"
	emailerPB "github.com/xmlking/micro-starter-kit/srv/emailer/proto/emailer"
	"github.com/xmlking/micro-starter-kit/srv/emailer/service"
)

// EmailSubscriber is Subscriber
type EmailSubscriber struct {
	emailService service.EmailService
}

// NewEmailSubscriber returns an instance of `EmailSubscriber`.
func NewEmailSubscriber(emailService service.EmailService) *EmailSubscriber {
	return &EmailSubscriber{
		emailService: emailService,
	}
}

// Handle is a method to send emails
func (s *EmailSubscriber) Handle(ctx context.Context, msg *emailerPB.Message) error {
	log.Infof("Sending email to: %s from: %s subject: %s body: %s", msg.To, msg.From, msg.Subject, msg.Body)
	s.emailService.Welcome(msg.Subject, msg.To)
	return nil
}

// Handler is a function to send emails
func Handler(ctx context.Context, msg *emailerPB.Message) error {
	log.Infof("Sending email to: %s from: %s subject: %s body: %s", msg.To, msg.From, msg.Subject, msg.Body)
	// TODO delegate to emailService.Welcome
	return nil
}
