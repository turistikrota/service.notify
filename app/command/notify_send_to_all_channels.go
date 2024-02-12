package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
)

type NotifySendToAllChannelsCmd struct{}

type NotifySendToAllChannelsRes struct{}

type NotifySendToAllChannelsHandler cqrs.HandlerFunc[NotifySendToAllChannelsCmd, *NotifySendToAllChannelsRes]

func NewNotifySendToAllChannelsHandler() NotifySendToAllChannelsHandler {
	return func(ctx context.Context, cmd NotifySendToAllChannelsCmd) (*NotifySendToAllChannelsRes, *i18np.Error) {
		return &NotifySendToAllChannelsRes{}, nil
	}
}
