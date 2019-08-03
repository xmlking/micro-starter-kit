package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/email"
)

type FakeEmailSender struct {
	mock.Mock
}

func (mock *FakeEmailSender) Send(subject, body string, to []string) error {
	args := mock.Called(subject, body, to)
	return args.Error(0)
}
func TestEmailService_Welcome(t *testing.T) {
	myConfig.InitConfig("../../../config", "config.yaml")
	emailer := &FakeEmailSender{}
	emailer.On("Send",
		"Welcome", "Hi Bob!", []string{"bob@smith.com"}).Return(nil)

	welcomer := NewEmailService(emailer)

	err := welcomer.Welcome("Bob", "bob@smith.com")
	assert.NoError(t, err)
	emailer.AssertExpectations(t)
}

func TestEmailService_Welcome_Integration(t *testing.T) {
	var (
		cfg myConfig.ServiceConfiguration
	)

	if testing.Short() {
		t.Skip("skipping long integration test")
	}
	myConfig.InitConfig("../../../config", "config.test.yaml")

	emailer := email.NewSendEmail(&cfg.Email)
	emailService := NewEmailService(emailer)
	t.Logf("emailer: %v", emailer)

	err2 := emailService.Welcome("Welcome", "sumo")
	if err2 != nil {
		t.Errorf("Send Welcome Email Failed: %v", err2)
	}
}
