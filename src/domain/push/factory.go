package push

import (
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/validator"
	"github.com/turistikrota/service.notify/src/domain/notify"
)

type Factory struct {
	Errors        Errors
	validator     *validator.Validator
	notifyFactory notify.Factory
}

func NewFactory(val *validator.Validator, notifyFactory notify.Factory) Factory {
	return Factory{
		Errors:        newPushErrors(),
		validator:     val,
		notifyFactory: notifyFactory,
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

func (f Factory) NewNotifyPush(token string, notification *Notification) (*Entity, *i18np.Error) {
	n, err := f.notifyFactory.NewNotify(token, notify.ChannelPush, f.notifyFactory.MarshalData(notification))
	if err != nil {
		return nil, err
	}
	t := time.Now()
	time := t.Format("04:05")
	return &Entity{
		Notify: n,
		Token:  token,
		Data: &Data{
			Score: "850",
			Time:  time,
		},
		Notification: notification,
	}, nil
}

func (f Factory) GetDataAsPush(n *notify.Notify) (*Entity, *i18np.Error) {
	if n.Type != notify.ChannelPush {
		return nil, f.notifyFactory.Errors.TypeNotFound("Push")
	}
	var e Entity
	if err := f.notifyFactory.ParseData(n.Data, &e); err != nil {
		return nil, f.Errors.Failed("parse entity")
	}
	return &e, nil
}

func (f Factory) Validate(n *notify.Notify) *i18np.Error {
	sms, err := f.GetDataAsPush(n)
	if err != nil {
		return err
	}
	if err := f.validator.ValidateStruct(sms); err != nil {
		return f.notifyFactory.Errors.ValidationFailed("create push")
	}
	return nil
}
