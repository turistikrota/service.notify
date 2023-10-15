package notify

import (
	"encoding/json"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/validator"
)

type Factory struct {
	Errors    Errors
	validator *validator.Validator
	debug     bool
}

func NewFactory(val *validator.Validator, debug bool) Factory {
	return Factory{
		Errors:    newNotifyErrors(),
		validator: val,
		debug:     debug,
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

func (f Factory) createTime() time.Time {
	if f.debug {
		return time.Time{}
	}
	return time.Now()
}

func (f Factory) NewNotify(recipient string, ch Channel, data []byte) (*Notify, *i18np.Error) {
	t := f.createTime()
	n := &Notify{
		UUID:      "",
		Type:      ch,
		Recipient: recipient,
		CreatedAt: t,
		UpdatedAt: t,
		Data:      data,
	}
	if err := f.Validate(n); err != nil {
		return nil, err
	}
	return n, nil
}

func (f Factory) Validate(n *Notify) *i18np.Error {
	return nil
}

func (f Factory) Unmarshal(uuid string, data []byte) *Notify {
	var n Notify
	if err := json.Unmarshal(data, &n); err != nil {
		return nil
	}
	n.UUID = uuid
	return &n
}

func (f Factory) MarshalData(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}

func (f Factory) ParseData(n []byte, data interface{}) error {
	return json.Unmarshal(n, data)
}
