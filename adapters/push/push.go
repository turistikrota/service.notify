package push

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/turistikrota/service.notify/config"
	"google.golang.org/api/option"
)

type SendConfig struct {
	Token string `json:"token"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Image string `json:"image"`
}

type SendAllConfig struct {
	Tokens []string `json:"tokens"`
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Image  string   `json:"image"`
}

type Service interface {
	Send(ctx context.Context, cnf SendConfig) error
	SendAll(ctx context.Context, cnf SendAllConfig) error
}

type srv struct {
	client *messaging.Client
}

func New(cnf config.Firebase) Service {
	opt := option.WithCredentialsFile(cnf.SecretFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(fmt.Errorf("error initializing firebase app:%v", err))
	}
	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error getting firebase Messaging client: %v\n", err)
	}
	return &srv{
		client: client,
	}
}

func (s *srv) Send(ctx context.Context, cnf SendConfig) error {
	t := time.Now()
	time := t.Format("04:05")
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  time,
		},
		Notification: &messaging.Notification{
			ImageURL: cnf.Image,
			Title:    cnf.Title,
			Body:     cnf.Body,
		},
		Token: cnf.Token,
	}
	_, err := s.client.Send(ctx, message)
	if err != nil {
		return err
	}
	return nil
}

func (s *srv) SendAll(ctx context.Context, cnf SendAllConfig) error {
	t := time.Now()
	time := t.Format("04:05")
	message := &messaging.MulticastMessage{
		Data: map[string]string{
			"score": "850",
			"time":  time,
		},
		Notification: &messaging.Notification{
			ImageURL: cnf.Image,
			Title:    cnf.Title,
			Body:     cnf.Body,
		},
		Tokens: cnf.Tokens,
	}

	br, err := s.client.SendMulticast(ctx, message)
	if err != nil {
		return err
	}
	if br.FailureCount > 0 {
		return errors.New("failed to send push")
	}
	return nil
}
