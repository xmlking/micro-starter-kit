package email

import (
	"bytes"
	"html/template"
	"net/smtp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/xmlking/micro-starter-kit/shared/config"

	myErrors "github.com/xmlking/micro-starter-kit/shared/errors"
)

var emailTmpl *template.Template

func init() {
	tmpl := `From: {{.From}}<br />
    To: {{.To}}<br />
    Subject: {{.Subject}}<br />
    MIME-version: 1.0<br />
    Content-Type: text/html; charset=&quot;UTF-8&quot;<br />
    <br />
    {{.Message}}`

	emailTmpl = template.Must(template.New("emailTemplate").Parse(tmpl))
}

// SendEmail struct
type SendEmail struct {
	from    string
	address string
	auth    smtp.Auth
}

// Send sends an email here, and perhaps returns an error.
func (sender *SendEmail) Send(subject, body string, to []string) error {
	var doc bytes.Buffer
	context := struct {
		From    string
		To      string
		Subject string
		Message string
	}{
		sender.from,
		strings.Join([]string(to), ","),
		subject,
		body,
	}
	err1 := emailTmpl.Execute(&doc, context)
	if err1 != nil {
		log.Print("error trying to execute mail template")
		return err1
	}
	log.Debugf("sending email to: %s from: %s, subject: %s, body: %s", to, sender.from, subject, body)
	err := smtp.SendMail(sender.address, sender.auth, sender.from, to, doc.Bytes())
	if err != nil {
		return myErrors.AppError(myErrors.SME, err.Error())
	}
	return nil
}

// NewSendEmail is constructor
func NewSendEmail(emailConf *config.EmailConfiguration) *SendEmail {
	return &SendEmail{
		from:    emailConf.From,
		address: emailConf.EmailServer + ":" + strconv.Itoa(emailConf.Port),
		auth:    smtp.PlainAuth("", emailConf.Username, emailConf.Password, emailConf.EmailServer),
	}
}
