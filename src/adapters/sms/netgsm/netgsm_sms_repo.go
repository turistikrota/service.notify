package netgsm

import (
	"api.turistikrota.com/notify/src/config"
	"api.turistikrota.com/notify/src/domain/notify"
	"api.turistikrota.com/notify/src/domain/sms"
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
