package sms

import (
	"api.turistikrota.com/notify/src/adapters/sms/netgsm"
	"api.turistikrota.com/notify/src/config"
	"api.turistikrota.com/notify/src/domain/notify"
	"api.turistikrota.com/notify/src/domain/sms"
)

type Sms interface {
	NewNetGSM(factory sms.Factory, notifyFactory notify.Factory, config config.NetGsm) sms.Repository
}

type smsClient struct{}

func New() Sms {
	return &smsClient{}
}

func (s *smsClient) NewNetGSM(factory sms.Factory, notifyFactory notify.Factory, config config.NetGsm) sms.Repository {
	return netgsm.New(factory, notifyFactory, config)
}
