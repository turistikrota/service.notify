package actor_config

import (
	"time"
)

type Entity struct {
	UUID      string               `json:"uuid" bson:"_id,omitempty"`
	Actor     Actor                `json:"actor" bson:"actor"`
	Telegram  []TelegramCredential `json:"telegram" bson:"telegram"`
	Mail      []MailCredential     `json:"mail" bson:"mail"`
	SMS       []SMSCredential      `json:"sms" bson:"sms"`
	UpdatedAt time.Time            `json:"updatedAt" bson:"updated_at"`
}

type TelegramCredential struct {
	Name   string `json:"name" bson:"name"`
	ChatID string `json:"chatId" bson:"chat_id"`
}

type MailCredential struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}

type SMSCredential struct {
	Name        string `json:"name" bson:"name"`
	Phone       string `json:"phone" bson:"phone" validate:"required,min=8,max=15"`
	CountryCode string `json:"countryCode" bson:"country_code" validate:"required,min=1,max=5"`
}

type Actor struct {
	UUID string    `json:"uuid" bson:"uuid"`
	Name string    `json:"name" bson:"name"`
	Type ActorType `json:"type" bson:"type"`
}

type ActorType string
type Type string

const (
	ActorTypeUser     ActorType = "user"
	ActorTypeBusiness ActorType = "business"

	TypeMail     Type = "mail"
	TypeSMS      Type = "sms"
	TypeTelegram Type = "telegram"
)

func (a ActorType) String() string {
	return string(a)
}

func (t Type) String() string {
	return string(t)
}
