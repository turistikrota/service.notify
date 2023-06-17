package telegram

import "api.turistikrota.com/notify/src/domain/notify"

type Data struct {
	ChatID string `json:"chat_id" validate:"required,min=1,max=255" bson:"chat_id"`
	Text   string `json:"text" validate:"required,min=1,max=65535" bson:"text"`
}

type Telegram struct {
	Data   *Data
	Notify *notify.Notify
}
