package dto

import (
	"github.com/turistikrota/service.notify/src/app/command"
	"github.com/turistikrota/service.notify/src/domain/push"
)

type NotifyPush struct {
	push.Notification
	DeviceUUID string `json:"device_uuid"`
	UserUUID   string `json:"user_uuid"`
}

func (n *NotifyPush) ToCommand() command.SendPushCommand {
	return command.SendPushCommand{
		DeviceUUID:   n.DeviceUUID,
		Notification: &n.Notification,
		UserUUID:     n.UserUUID,
	}
}
