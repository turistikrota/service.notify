package notify

import (
	"github.com/turistikrota/service.notify/src/domain/notify"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	notifyFactory notify.Factory
	collection    *mongo.Collection
}

func New(notifyFactory notify.Factory, collection *mongo.Collection) notify.Repository {
	validate(notifyFactory, collection)
	return &repo{
		notifyFactory: notifyFactory,
		collection:    collection,
	}
}

func validate(notifyFactory notify.Factory, collection *mongo.Collection) {
	if notifyFactory.IsZero() {
		panic("notifyFactory is zero")
	}
	if collection == nil {
		panic("collection is nil")
	}
}
