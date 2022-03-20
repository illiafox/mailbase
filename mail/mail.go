package mail

import (
	"gopkg.in/gomail.v2"
	"mailbase/util/config"
)

func NewMail(conf config.Config) Mail {
	// Connect to the remote SMTP server.
	return Mail{
		Dialer: gomail.NewDialer(conf.Smtp.Hostname, conf.Smtp.Port, conf.Smtp.Mail, conf.Smtp.Password),
		From:   conf.Smtp.Mail,
	}
}
