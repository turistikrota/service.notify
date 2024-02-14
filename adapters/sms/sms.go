package sms

import (
	"bytes"
	"context"
	"net/http"

	"github.com/turistikrota/service.notify/config"
)

type SendConfig struct {
	Phone string `json:"phone"`
	Text  string `json:"text"`
	Lang  string `json:"lang"`
}

type Service interface {
	Send(ctx context.Context, config SendConfig) error
}

type srv struct {
	cnf config.NetGsm
}

func New(cnf config.NetGsm) Service {
	return &srv{
		cnf: cnf,
	}
}

func (s *srv) Send(ctx context.Context, config SendConfig) error {
	phone := config.Phone
	// remove + sign
	if phone[0] == '+' {
		phone = phone[1:]
	}
	msg := `<?xml version="1.0" encoding="utf-8"?>
    <mainbody>
        <header>
            <company dil="` + config.Lang + `">` + s.cnf.Company + `</company>
            <usercode>` + s.cnf.UserName + `</usercode>
            <password>` + s.cnf.Password + `</password>
            <type>1:n</type>
            <msgheader>` + s.cnf.Title + `</msgheader>
        </header>
        <body>
            <msg><![CDATA[` + config.Text + `]]></msg>
            <no>` + phone + `</no>
        </body>
    </mainbody>
    `
	req, err := http.NewRequest("POST", s.getApiUrl(), bytes.NewBuffer([]byte(msg)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/xml")
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *srv) getApiUrl() string {
	return "https://api.netgsm.com.tr/sms/send/xml"
}
