package netgsm

import (
	"bytes"
	"context"
	"net/http"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.notify/src/domain/sms"
)

func (r *repo) Send(ctx context.Context, config sms.SendConfig) *i18np.Error {
	msg := r.buildMsg(config.Data)
	req, err := http.NewRequest("POST", r.getApiUrl(), bytes.NewBuffer(msg))
	if err != nil {
		return r.factory.Errors.Failed(err.Error())
	}
	req.Header.Set("Content-Type", "application/xml")
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return r.factory.Errors.Failed(err.Error())
	}
	return nil
}

func (r *repo) buildMsg(data *sms.Data) []byte {
	phone := data.Phone
	// remove + sign
	if phone[0] == '+' {
		phone = phone[1:]
	}
	msg := `<?xml version="1.0" encoding="utf-8"?>
    <mainbody>
        <header>
            <company dil="` + data.Lang + `">` + r.cnf.Company + `</company>
            <usercode>` + r.cnf.UserName + `</usercode>
            <password>` + r.cnf.Password + `</password>
            <type>1:n</type>
            <msgheader>` + r.cnf.Title + `</msgheader>
        </header>
        <body>
            <msg><![CDATA[` + data.Text + `]]></msg>
            <no>` + phone + `</no>
        </body>
    </mainbody>
    `
	return []byte(msg)
}

func (r *repo) getApiUrl() string {
	return "https://api.netgsm.com.tr/sms/send/xml"
}
