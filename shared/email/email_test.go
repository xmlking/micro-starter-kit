package email

import (
	"fmt"
	"net/smtp"
	"strconv"
	"testing"

    "github.com/xmlking/micro-starter-kit/shared/config"
)

func TestSendEmail_Send(t *testing.T) {
	var emailConf = config.GetConfig().Email
	myAuth := smtp.PlainAuth("", emailConf.Username, emailConf.Password, emailConf.EmailServer)
	from := emailConf.From
	address := emailConf.EmailServer + ":" + strconv.FormatUint(uint64(emailConf.Port), 10)

	type fields struct {
		from    string
		address string
		auth    smtp.Auth
	}
	type args struct {
		subject string
		body    string
		to      []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "does not return an error when send method is success",
			fields:  fields{from: from, address: address, auth: myAuth},
			args:    args{subject: "subject1", body: "body1", to: []string{"bob@smith.com"}},
			wantErr: false,
		},
		{
			name:    "returns an error when input is invalid and send method fail",
			fields:  fields{from: "xyz@gmail.com", address: address, auth: myAuth},
			args:    args{subject: "subject2", body: "body2", to: []string{"bob2@smith.com"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sender := &SendEmail{
				from:    tt.fields.from,
				address: tt.fields.address,
				auth:    tt.fields.auth,
				// ADD Mock SendMail func
				send: func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
					t.Logf("from: %s, to: %s, msg: %s, address: %s, auth: %v", from, to, string(msg), addr, a)
					if from == "xyz@gmail.com" {
						return fmt.Errorf("we don't like %s", from)
					}
					return nil
				},
			}
			if err := sender.Send(tt.args.subject, tt.args.body, tt.args.to); (err != nil) != tt.wantErr {
				// expect myErrors.AppError(myErrors.SME, err)
				t.Errorf("SendEmail.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
