package telegram

import (
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.notify/src/domain/notify"
	"github.com/turistikrota/service.shared/validator"
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

func (f Factory) NewNotifyTelegram(recipient string, data *Data) (*Telegram, *i18np.Error) {
	n, err := f.notifyFactory.NewNotify(recipient, notify.ChannelTelegram, f.notifyFactory.MarshalData(data))
	if err != nil {
		return nil, err
	}
	return &Telegram{
		Notify: n,
		Data:   data,
	}, nil
}

func (f Factory) GetDataAsTelegram(n *notify.Notify) (*Data, *i18np.Error) {
	if n.Type != notify.ChannelTelegram {
		return nil, f.notifyFactory.Errors.TypeNotFound("Telegram")
	}
	var data Data
	if err := f.notifyFactory.ParseData(n.Data, &data); err != nil {
		return nil, f.Errors.Failed("get Telegram")
	}
	return &data, nil
}

func (f Factory) Validate(n *notify.Notify) *i18np.Error {
	telegram, err := f.GetDataAsTelegram(n)
	if err != nil {
		return err
	}
	if err := f.validator.ValidateStruct(telegram); err != nil {
		return f.notifyFactory.Errors.ValidationFailed("create Telegram")
	}
	return nil
}
