package push

import (
	firebaseGo "firebase.google.com/go"
	"github.com/turistikrota/service.notify/src/adapters/push/firebase"
	"github.com/turistikrota/service.notify/src/domain/notify"
	"github.com/turistikrota/service.notify/src/domain/push"
)

type Push interface {
	NewPush(factory push.Factory, notifyFactory notify.Factory, app *firebaseGo.App) push.Repository
}

type pushClient struct{}

func New() Push {
	return &pushClient{}
}

func (m *pushClient) NewPush(factory push.Factory, notifyFactory notify.Factory, app *firebaseGo.App) push.Repository {
	return firebase.New(factory, notifyFactory, app)
}
