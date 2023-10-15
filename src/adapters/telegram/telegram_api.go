package telegram

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.notify/src/domain/telegram"
)

func (r *repo) Send(ctx context.Context, config telegram.SendConfig) *i18np.Error {
	msg, err := r.formatMsgAsHTML(config.Data.Text, config.Data.ChatID)
	if err != nil {
		return r.factory.Errors.Failed(err.Error())
	}
	req, err := http.NewRequest("POST", r.getApiUrl(), bytes.NewBuffer(msg))
	if err != nil {
		return r.factory.Errors.Failed(err.Error())
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req.Header.Set("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return r.factory.Errors.Failed(err.Error())
	}
	return nil
}

func (r *repo) formatMsgAsHTML(text string, chatId string) ([]byte, error) {
	mp := map[string]interface{}{
		"chat_id":    chatId,
		"text":       text,
		"parse_mode": "HTML",
	}
	bytes, err := json.Marshal(mp)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (r *repo) getApiUrl() string {
	return "https://api.telegram.org/bot" + r.conf.Token + "/sendMessage"
}
