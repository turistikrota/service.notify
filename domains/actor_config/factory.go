package actor_config

import "time"

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

func (f Factory) New(actor Actor) *Entity {
	return &Entity{
		Actor:     actor,
		Telegram:  []TelegramCredential{},
		Mail:      []MailCredential{},
		SMS:       []SMSCredential{},
		UpdatedAt: time.Now(),
	}
}
