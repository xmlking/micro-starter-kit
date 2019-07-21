package registry

import (
	"github.com/sarulabs/di"

	log "github.com/sirupsen/logrus"
	"github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/email"
	emailerService "github.com/xmlking/micro-starter-kit/srv/emailer/service"
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
				return email.CreateSendEmail(&cfg.Email), nil
			},
		},
		{
			Name:  "email-service",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				emailer := ctn.Get("send-email").(email.SendEmail)
				return emailerService.CreateEmailService(emailer), nil
			},
		},
		{ // TODO
			Name:  "emailer-subscriber",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				emailer := ctn.Get("send-email").(email.SendEmail)
				return emailerService.CreateEmailService(emailer), nil
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
