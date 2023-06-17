package notify

import (
	"time"
)

type Channel string

type Notify struct {
	UUID      string
	Type      Channel
	Recipient string
	CreatedAt time.Time
	UpdatedAt time.Time
	Data      []byte
}

const (
	ChannelEmail    Channel = "email"
	ChannelSMS      Channel = "sms"
	ChannelTelegram Channel = "telegram"
)
