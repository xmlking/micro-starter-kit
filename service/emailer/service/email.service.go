package service

import (
    "bytes"
    "html/template"

    "github.com/rs/zerolog/log"
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
type emailService struct {
    Emailer EmailSender
}

// Welcome method
func (welcomer *emailService) Welcome(name, email string) error {
    log.Info().Msg("in Welcome")
    var body bytes.Buffer
    if err := welcomeEmailTmpl.Execute(&body, struct{ Name string }{name}); err != nil {
        return err
    }
    subject := "Welcome"

    return welcomer.Emailer.Send(subject, body.String(), []string{email})
}

// EmailService interface
type EmailService interface {
    Welcome(name, email string) error
}

// NewEmailService is constructor
func NewEmailService(Emailer EmailSender) EmailService {
    return &emailService{
        Emailer: Emailer,
    }
}
