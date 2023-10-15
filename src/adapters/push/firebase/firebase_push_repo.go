package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/turistikrota/service.notify/src/domain/notify"
	"github.com/turistikrota/service.notify/src/domain/push"
)

type repo struct {
	client        *messaging.Client
	factory       push.Factory
	notifyFactory notify.Factory
}

func New(factory push.Factory, notifyFactory notify.Factory, app *firebase.App) push.Repository {
	if notifyFactory.IsZero() {
		panic("notifyFactory is zero")
	}
	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	r := &repo{
		client:        client,
		factory:       factory,
		notifyFactory: notifyFactory,
	}
	return r
}
