package service

import (
	"bytes"
	"fmt"
	"net/smtp"
)

type mailer struct {
	Host     string
	Port     string
	User     string
	Password string
}

// Send an email to a user
func (m *mailer) Send(from, to, body string) {
	msg := bytes.NewBufferString(body)
	tos := []string{to}

	auth := smtp.PlainAuth("", m.User, m.Password, m.Host)
	err := smtp.SendMail(m.address(), auth, from, tos, msg.Bytes())
	if err != nil {
		Logger.Info("%s", err)
	}
}

func (m *mailer) address() string {
	return fmt.Sprintf("%s:%s", m.Host, m.Port)
}
