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
	Send(to, subject, body string) error
}

// EmailService struct
type EmailService struct {
	Emailer EmailSender
}

// CreateEmailService method
func CreateEmailService(Emailer EmailSender) *EmailService {
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

	return welcomer.Emailer.Send(email, subject, body.String())
}
