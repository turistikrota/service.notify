package mongo

import (
	notify_mongo "api.turistikrota.com/notify/src/adapters/mongo/notify"
	"api.turistikrota.com/notify/src/domain/notify"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo interface {
	NewNotify(notifyFactory notify.Factory, collection *mongo.Collection) notify.Repository
}

type mongoClient struct{}

func New() Mongo {
	return &mongoClient{}
}

func (m *mongoClient) NewNotify(notifyFactory notify.Factory, collection *mongo.Collection) notify.Repository {
	return notify_mongo.New(notifyFactory, collection)
}
