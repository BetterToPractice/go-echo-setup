package lib

import (
	"bytes"
	"fmt"
	mail "github.com/xhit/go-simple-mail/v2"
	"html/template"
	"os"
)

type Mail struct {
	smtp   *mail.SMTPClient
	config *MailConfig
	logger Logger
}

type MailTemplate struct {
	Subject   string
	Body      string
	Sender    string
	Receivers []string
	Context   map[string]interface{}
}

func NewMail(config Config, logger Logger) Mail {
	server := mail.NewSMTPClient()

	server.Host = config.Mail.Host
	server.Port = config.Mail.Port
	server.Username = config.Mail.User
	server.Password = config.Mail.Password
	if config.Mail.UseTLS {
		server.Encryption = mail.EncryptionSTARTTLS
	}
	server.KeepAlive = true

	smtp, err := server.Connect()
	if err != nil && config.Mail.Enable {
		logger.Zap.Fatalf("Error to open mail [%s] connection: %v", config.Mail.Host, err)
	}

	return Mail{smtp: smtp, config: config.Mail, logger: logger}
}

func (l Mail) SendMailWithTemplate(mailTemplate MailTemplate) {
	sender := l.config.FromEmail
	if mailTemplate.Sender != "" {
		sender = mailTemplate.Sender
	}

	l.SendMail(mailTemplate.GetSubject(), mailTemplate.GetBody(), mailTemplate.Receivers, sender)
}

func (l Mail) SendMailAsLog(subject string, body string, receivers []string, sender string) {
	fmt.Printf(
		"=====================================================================================\n"+
			"[Subject]:\t%s\n[Sender]:\t%s\n[Receivers]:\t%s\n[Body]\n%s\n"+
			"=====================================================================================\n",
		subject,
		sender,
		receivers,
		body,
	)
}

func (l Mail) SendMail(subject string, body string, receivers []string, sender string) {
	if !l.config.Enable {
		l.SendMailAsLog(subject, body, receivers, sender)
		return
	}

	for _, to := range receivers {
		email := mail.NewMSG()
		if sender != "" {
			email.SetFrom(sender).AddTo(to).SetSubject(subject).SetBody(mail.TextHTML, body)
		} else {
			email.SetFrom(l.config.FromEmail).AddTo(to).SetSubject(subject).SetBody(mail.TextHTML, body)
		}

		if email.Error != nil {
			l.logger.DesugarZap.Warn(email.Error.Error())
			return
		}

		err := email.Send(l.smtp)
		if err != nil {
			l.logger.DesugarZap.Warn(email.Error.Error())
			return
		}
	}
}

func (m MailTemplate) GetSubject() string {
	tmpl, err := m.ReadTemplate(m.Subject)
	if err != nil {
		return ""
	}
	return tmpl
}

func (m MailTemplate) GetBody() string {
	tmpl, err := m.ReadTemplate(m.Body)
	if err != nil {
		return ""
	}
	return tmpl
}

func (m MailTemplate) ReadTemplate(templateName string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	templateContent, err := os.ReadFile(
		fmt.Sprintf("%s/templates/%s", cwd, templateName),
	)
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("emailTemplate").Parse(string(templateContent))
	if err != nil {
		return "", err
	}

	var emailBodyBuffer bytes.Buffer
	if err := tmpl.Execute(&emailBodyBuffer, m.Context); err != nil {
		return "", err
	}
	return emailBodyBuffer.String(), nil
}
