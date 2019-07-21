package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
)

type FakeEmailSender struct {
	mock.Mock
}

func (mock *FakeEmailSender) Send(to, subject, body string) error {
	args := mock.Called(to, subject, body)
	return args.Error(0)
}
func TestEmailService_Welcome(t *testing.T) {
	myConfig.InitConfig("../../../config", "config.yaml")

	emailer := &FakeEmailSender{}
	emailer.On("Send",
		"bob@smith.com", "Welcome", "Hi Bob!").Return(nil)

	welcomer := CreateEmailService(emailer)

	err := welcomer.Welcome("Bob", "bob@smith.com")
	assert.NoError(t, err)
	emailer.AssertExpectations(t)
}
