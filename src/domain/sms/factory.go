package sms

import (
	"api.turistikrota.com/notify/src/domain/notify"
	"github.com/turistikrota/service.shared/validator"
	"github.com/mixarchitecture/i18np"
)

type Factory struct {
	Errors        Errors
	validator     *validator.Validator
	notifyFactory notify.Factory
}

func NewFactory(val *validator.Validator, notifyFactory notify.Factory) Factory {
	return Factory{
		Errors:        newSmsErrors(),
		validator:     val,
		notifyFactory: notifyFactory,
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

func (f Factory) NewNotifySms(recipient string, data *Data) (*Sms, *i18np.Error) {
	n, err := f.notifyFactory.NewNotify(recipient, notify.ChannelSMS, f.notifyFactory.MarshalData(data))
	if err != nil {
		return nil, err
	}
	return &Sms{
		Notify: n,
		Data:   data,
	}, nil
}

func (f Factory) GetDataAsSMS(n *notify.Notify) (*Data, *i18np.Error) {
	if n.Type != notify.ChannelSMS {
		return nil, f.notifyFactory.Errors.TypeNotFound("SMS")
	}
	var data Data
	if err := f.notifyFactory.ParseData(n.Data, &data); err != nil {
		return nil, f.Errors.Failed("get SMS")
	}
	return &data, nil
}

func (f Factory) Validate(n *notify.Notify) *i18np.Error {
	sms, err := f.GetDataAsSMS(n)
	if err != nil {
		return err
	}
	if err := f.validator.ValidateStruct(sms); err != nil {
		return f.notifyFactory.Errors.ValidationFailed("create SMS")
	}
	return nil
}
