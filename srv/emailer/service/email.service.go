package service

import (
	"bytes"
	"html/template"
)

var welcomeEmailTmpl *template.Template

func init() {
	tmpl := `Hi {{.Name}}!`

	welcomeEmailTmpl = template.Must(template.New("welcomeEmailTemplate").Parse(tmpl))
}

// EmailSender provides an interface so we can swap out the
// implementation of SendEmail under tests.
type EmailSender interface {
	Send(subject, body string, to []string) error
}

// EmailService struct
type EmailService struct {
	Emailer EmailSender
}

// NewEmailService method
func NewEmailService(Emailer EmailSender) *EmailService {
	return &EmailService{
		Emailer: Emailer,
	}
}

// Welcome method
func (welcomer *EmailService) Welcome(name, email string) error {
	var body bytes.Buffer
	if err := welcomeEmailTmpl.Execute(&body, struct{ Name string }{name}); err != nil {
		return err
	}
	subject := "Welcome"

	return welcomer.Emailer.Send(subject, body.String(), []string{email})
}
