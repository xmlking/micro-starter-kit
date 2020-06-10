package subscriber

import (
    "context"

    "github.com/micro/go-micro/v2/metadata"
    "github.com/rs/zerolog/log"

    emailerPB "github.com/xmlking/micro-starter-kit/service/emailer/proto/emailer"
    "github.com/xmlking/micro-starter-kit/service/emailer/service"
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

// Handle is a method to send emails, Method can be of any name
func (s *EmailSubscriber) Handle(ctx context.Context, msg *emailerPB.Message) error {
    md, _ := metadata.FromContext(ctx)
    log.Info().Msgf("EmailSubscriber: Received msg %+v with metadata %+v\n", msg, md)
    return s.emailService.Welcome(msg.Subject, msg.To)
}

// Handler is a function to send emails
func Handler(ctx context.Context, msg *emailerPB.Message) error {
    md, _ := metadata.FromContext(ctx)
    log.Info().Msgf("Function: Received msg %+v with metadata %+v\n", msg, md)
    // TODO delegate to emailService.Welcome
    return nil
}
