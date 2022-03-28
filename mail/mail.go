package mail

import (
	"fmt"
	"github.com/illiafox/mailbase/util/config"
	"gopkg.in/gomail.v2"
)

func NewMail(conf config.Config) (Mail, error) {
	m := Mail{
		Dialer: gomail.NewDialer(conf.SMTP.Hostname, conf.SMTP.Port, conf.SMTP.Mail, conf.SMTP.Password),
		From:   conf.SMTP.Mail,
	}

	// Connect to the remote SMTP server.
	rc, err := m.Dialer.Dial()
	if err != nil {
		return Mail{}, fmt.Errorf("dialing: %w", err)
	}

	err = rc.Close()
	if err != nil {
		return Mail{}, fmt.Errorf("closing test dial: %w", err)
	}

	return m, nil
}
