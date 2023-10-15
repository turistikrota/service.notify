package mail

import (
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
		Errors:        newSmsErrors(),
		validator:     val,
		notifyFactory: notifyFactory,
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

func (f Factory) NewNotifyMail(recipient string, data *Data) (*Mail, *i18np.Error) {
	n, err := f.notifyFactory.NewNotify(recipient, notify.ChannelEmail, f.notifyFactory.MarshalData(data))
	if err != nil {
		return nil, err
	}
	return &Mail{
		Notify: n,
		Data:   data,
	}, nil
}

func (f Factory) GetDataAsMail(n *notify.Notify) (*Data, *i18np.Error) {
	if n.Type != notify.ChannelEmail {
		return nil, f.notifyFactory.Errors.TypeNotFound("Email")
	}
	var data Data
	if err := f.notifyFactory.ParseData(n.Data, &data); err != nil {
		return nil, f.Errors.Failed("get Email")
	}
	return &data, nil
}

func (f Factory) Validate(n *notify.Notify) *i18np.Error {
	mail, err := f.GetDataAsMail(n)
	if err != nil {
		return err
	}
	if err := f.validator.ValidateStruct(mail); err != nil {
		return f.notifyFactory.Errors.ValidationFailed("create Email")
	}
	return nil
}
