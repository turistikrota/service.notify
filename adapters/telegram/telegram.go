package telegram

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/turistikrota/service.notify/config"
)

type SendConfig struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

type Service interface {
	Send(ctx context.Context, config SendConfig) error
}

type srv struct {
	cnf config.Telegram
}

func New(cnf config.Telegram) Service {
	return &srv{
		cnf: cnf,
	}
}

func (s *srv) Send(ctx context.Context, config SendConfig) error {
	msg, err := s.formatMsgAsHTML(config.Text, config.ChatID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req, err := http.NewRequest("POST", s.getApiUrl(), bytes.NewBuffer(msg))
	if err != nil {
		fmt.Println(err)
		return err
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req.Header.Set("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *srv) formatMsgAsHTML(text string, chatId string) ([]byte, error) {
	mp := map[string]interface{}{
		"chat_id":    chatId,
		"text":       text,
		"parse_mode": "HTML",
	}
	bytes, err := json.Marshal(mp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bytes, nil
}

func (s *srv) getApiUrl() string {
	return "https://api.telegram.org/bot" + s.cnf.Token + "/sendMessage"
}
