package sms

import "github.com/turistikrota/service.notify/src/domain/notify"

type Data struct {
	Phone string `json:"phone" validate:"required,phone" bson:"phone"`
	Text  string `json:"text" validate:"required,min=1,max=480" bson:"text"`
	Lang  string `json:"lang" validate:"required,min=1,max=2" bson:"lang_code"`
}

type Sms struct {
	Data   *Data
	Notify *notify.Notify
}
