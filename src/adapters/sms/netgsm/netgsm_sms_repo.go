package netgsm

import (
	"github.com/turistikrota/service.notify/src/config"
	"github.com/turistikrota/service.notify/src/domain/notify"
	"github.com/turistikrota/service.notify/src/domain/sms"
)

type repo struct {
	factory       sms.Factory
	notifyFactory notify.Factory
	cnf           config.NetGsm
}

func New(factory sms.Factory, notifyFactory notify.Factory, config config.NetGsm) sms.Repository {
	if notifyFactory.IsZero() {
		panic("notifyFactory is zero")
	}
	return &repo{
		factory:       factory,
		notifyFactory: notifyFactory,
		cnf:           config,
	}
}
