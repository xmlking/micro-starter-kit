package registry

import (
	"github.com/sarulabs/di/v2"

	log "github.com/sirupsen/logrus"
	"github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/email"
	"github.com/xmlking/micro-starter-kit/srv/emailer/service"
	"github.com/xmlking/micro-starter-kit/srv/emailer/subscriber"
)

// Container - provide di Container
type Container struct {
	ctn di.Container
}

// NewContainer - create new Container
func NewContainer(cfg config.ServiceConfiguration) (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name:  "config",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return cfg, nil
			},
		},
		{
			Name:  "send-email",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return email.NewSendEmail(&cfg.Email), nil
			},
		},
		{
			Name:  "email-service",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				emailer := ctn.Get("send-email").(email.SendEmail)
				return service.NewEmailService(&emailer), nil
			},
		},
		{
			Name:  "emailer-subscriber",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				emailService := ctn.Get("email-service").(service.EmailService)
				return subscriber.NewEmailSubscriber(emailService), nil
			},
		},
	}...); err != nil {
		return nil, err
	}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

// Resolve object
func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

// Clean Container
func (c *Container) Clean() error {
	return c.ctn.Clean()
}

// Delete Container
func (c *Container) Delete() error {
	return c.ctn.Delete()
}
