package sms

import (
	"github.com/turistikrota/service.notify/src/adapters/sms/netgsm"
	"github.com/turistikrota/service.notify/src/config"
	"github.com/turistikrota/service.notify/src/domain/notify"
	"github.com/turistikrota/service.notify/src/domain/sms"
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
