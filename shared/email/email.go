package email

import (
    "bytes"
    "html/template"
    "net/smtp"
    "strconv"
    "strings"

    "github.com/rs/zerolog/log"

    myErrors "github.com/xmlking/micro-starter-kit/shared/errors"
    configPB "github.com/xmlking/micro-starter-kit/shared/proto/config"
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
    send    func(addr string, a smtp.Auth, from string, to []string, msg []byte) error
}

// Send sends an email here, and perhaps returns an error.
func (sender *SendEmail) Send(subject, body string, to []string) error {
    log.Info().Msg("in SendEmail.Send")
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
        log.Error().Msg("error trying to execute mail template")
        return err1
    }
    log.Debug().Msgf("sending email to: %s from: %s, subject: %s, body: %s", to, sender.from, subject, doc.Bytes())
    err := sender.send(sender.address, sender.auth, sender.from, to, doc.Bytes())
    if err != nil {
        return myErrors.AppError(myErrors.SME, err.Error())
    }
    return nil
}

// NewSendEmail is constructor
func NewSendEmail(emailConf *configPB.EmailConfiguration) *SendEmail {
    return &SendEmail{
        from:    emailConf.From,
        address: emailConf.EmailServer + ":" + strconv.FormatUint(uint64(emailConf.Port), 10),
        auth:    smtp.PlainAuth("", emailConf.Username, emailConf.Password, emailConf.EmailServer),
        send:    smtp.SendMail,
    }
}
