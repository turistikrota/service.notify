package google

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.notify/src/domain/mail"
)

func (r *repo) Send(ctx context.Context, config mail.SendConfig) *i18np.Error {
	msg, err := r.dressUpTemplate(config.Data.Template, config.Data.Data)
	if err != nil {
		return r.factory.Errors.Failed(err.Error())
	}
	err = r.send([]string{config.Data.To}, r.buildMsg(config.Data.Subject, msg))
	if err != nil {
		return r.factory.Errors.Failed(err.Error())
	}
	return nil
}

func (r *repo) buildMsg(subject string, html string) []byte {
	subj := "Subject: " + subject + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	return []byte(subj + mime + html)
}

func (r *repo) send(to []string, msg []byte) error {
	return smtp.SendMail(r.getHost(), r.auth, r.conf.FromMail, to, msg)
}

func (r *repo) getHost() string {
	return r.conf.SmtpHost + ":" + r.conf.SmtpPort
}

func (r *repo) dressUpTemplate(temp string, data interface{}) (string, error) {
	temp = r.getTemplateUrl(temp)
	if _, err := os.Stat(temp); os.IsNotExist(err) {
		temp = r.getTemplateUrl("default")
	}
	t, err := template.ParseFiles(temp)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (r *repo) getTemplateUrl(temp string) string {
	return fmt.Sprintf("./assets/template/mail/%s.html", temp)
}
