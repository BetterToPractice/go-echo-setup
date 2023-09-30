package mails

import (
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/BetterToPractice/go-echo-setup/models"
)

type PostMail struct {
	mail lib.Mail
}

func NewPostMail(mail lib.Mail) PostMail {
	return PostMail{
		mail: mail,
	}
}

func (m PostMail) CreatePost(user *models.User, post *models.Post) {
	m.mail.SendMailWithTemplate(lib.MailTemplate{
		Subject:   "mails/post/post_create_subject.html",
		Body:      "mails/post/post_create_body.html",
		Receivers: []string{user.Email},
		Context: map[string]interface{}{
			"Username": user.Username,
			"Title":    post.Title,
		},
	})
}
