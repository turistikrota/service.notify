package mail

import "github.com/turistikrota/service.notify/src/domain/notify"

type Data struct {
	To       string      `json:"to" validate:"required,email" bson:"to"`
	Subject  string      `json:"subject" validate:"required,min=1,max=255" bson:"subject"`
	Data     interface{} `json:"data" validate:"required" bson:"data"`
	Template string      `json:"template" validate:"required,min=1,max=255" bson:"template"`
}

type Mail struct {
	Data   *Data
	Notify *notify.Notify
}
