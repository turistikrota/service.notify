package mail

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"github.com/turistikrota/service.notify/assets"
	"github.com/turistikrota/service.notify/config"
	smtp_mail "github.com/xhit/go-simple-mail/v2"
)

type SendConfig struct {
	To      string
	Subject string
	Message string
}

type SendWithTemplateConfig struct {
	SendConfig
	Template string
	Data     any
}

type Service interface {
	SendText(SendConfig) error
	SendWithTemplate(SendWithTemplateConfig) error
}

type srv struct {
	client *smtp_mail.SMTPClient
	config config.Smtp
}

func New(cnf config.Smtp) Service {
	server := smtp_mail.NewSMTPClient()
	server.Host = cnf.Host
	server.Port = cnf.Port
	server.Username = cnf.Sender
	server.Password = cnf.Password
	server.Encryption = smtp_mail.EncryptionSTARTTLS
	server.Authentication = smtp_mail.AuthLogin
	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}
	return &srv{
		client: smtpClient,
		config: cnf,
	}
}

func (s *srv) SendText(cnf SendConfig) error {
	email := smtp_mail.NewMSG()
	email.SetFrom(s.config.From)
	email.AddTo(cnf.To)
	email.SetSubject(cnf.Subject)
	email.SetSender(s.config.Sender)
	email.SetReplyTo(s.config.Reply)
	email.AddAlternative(smtp_mail.TextPlain, cnf.Message)
	err := email.Send(s.client)
	if err != nil {
		return err
	}
	return nil
}

func (s *srv) SendWithTemplate(cnf SendWithTemplateConfig) error {
	dir := assets.EmbedMailTemplate()
	t := template.Must(template.ParseFS(dir, fmt.Sprintf("mails/%s.html", cnf.Template)))
	var tpl bytes.Buffer
	t.Execute(&tpl, cnf.Data)
	body := tpl.String()
	email := smtp_mail.NewMSG()
	email.SetFrom(s.config.From)
	email.AddTo(cnf.To)
	email.SetSubject(cnf.Subject)
	email.SetSender(s.config.Sender)
	email.SetReplyTo(s.config.Reply)
	email.SetBody(smtp_mail.TextHTML, body)
	if err := email.Send(s.client); err != nil {
		return err
	}
	return nil
}
