package entity

import (
	"time"

	"api.turistikrota.com/notify/src/domain/notify"
)

type MongoNotify struct {
	UUID      string         `bson:"_id,omitempty"`
	Recipient string         `bson:"recipient"`
	Type      notify.Channel `bson:"type"`
	Data      interface{}    `bson:"data"`
	CreatedAt time.Time      `bson:"created_at"`
	UpdatedAt time.Time      `bson:"updated_at"`
}

func (m *MongoNotify) ToNotify() (*notify.Notify, interface{}) {
	return &notify.Notify{
		UUID:      m.UUID,
		Recipient: m.Recipient,
		Type:      m.Type,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}, m.Data
}

func (m *MongoNotify) ToDetail() *notify.DetailResult {
	return &notify.DetailResult{
		UUID:      m.UUID,
		Recipient: m.Recipient,
		Data:      m.Data,
		Type:      m.Type,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (m *MongoNotify) FromNotify(n *notify.Notify, data interface{}) *MongoNotify {
	if len(n.UUID) > 0 {
		m.UUID = n.UUID
	}
	m.Recipient = n.Recipient
	m.Type = n.Type
	m.Data = data
	m.CreatedAt = n.CreatedAt
	m.UpdatedAt = n.UpdatedAt
	return m
}
