package dto

import (
	"github.com/turistikrota/service.notify/src/app/command"
	"github.com/turistikrota/service.notify/src/domain/push"
)

type NotifyPush struct {
	push.Notification
	Token string `json:"token"`
}

func (n *NotifyPush) ToCommand() command.SendPushCommand {
	return command.SendPushCommand{
		Token:        n.Token,
		Notification: &n.Notification,
	}
}
