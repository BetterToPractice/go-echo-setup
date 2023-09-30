package mails

import (
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/BetterToPractice/go-echo-setup/models"
)

type AuthMail struct {
	mail lib.Mail
}

func NewAuthMail(mail lib.Mail) AuthMail {
	return AuthMail{
		mail: mail,
	}
}

func (m AuthMail) Register(user *models.User) {
	m.mail.SendMailWithTemplate(lib.MailTemplate{
		Subject:   "mails/auth/register_subject.html",
		Body:      "mails/auth/register_body.html",
		Receivers: []string{user.Email},
		Context:   map[string]interface{}{"Username": user.Username},
	})
}
