package mail

import (
	"bytes"
	"fmt"
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
	server *smtp_mail.SMTPServer
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
	return &srv{
		server: server,
		config: cnf,
	}
}

func (s *srv) createClient() (*smtp_mail.SMTPClient, error) {
	return s.server.Connect()
}

func (s *srv) SendText(cnf SendConfig) error {
	client, err := s.createClient()
	if err != nil {
		return err
	}
	email := smtp_mail.NewMSG()
	email.SetFrom(s.config.From)
	email.AddTo(cnf.To)
	email.SetSubject(cnf.Subject)
	email.SetSender(s.config.Sender)
	email.SetReplyTo(s.config.Reply)
	email.AddAlternative(smtp_mail.TextPlain, cnf.Message)
	err = email.Send(client)
	if err != nil {
		return err
	}
	return nil
}

func (s *srv) SendWithTemplate(cnf SendWithTemplateConfig) error {
	client, err := s.createClient()
	if err != nil {
		return err
	}
	dir := assets.EmbedMailTemplate()
	t := template.Must(template.ParseFS(dir, fmt.Sprintf("mail/%s.html", cnf.Template)))
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
	if err = email.Send(client); err != nil {
		return err
	}
	return nil
}
