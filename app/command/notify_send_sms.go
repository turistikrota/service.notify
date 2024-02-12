package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
)

type NotifySendSmsCmd struct{}

type NotifySendSmsRes struct{}

type NotifySendSmsHandler cqrs.HandlerFunc[NotifySendSmsCmd, *NotifySendSmsRes]

func NewNotifySendSmsHandler() NotifySendSmsHandler {
	return func(ctx context.Context, cmd NotifySendSmsCmd) (*NotifySendSmsRes, *i18np.Error) {
		return &NotifySendSmsRes{}, nil
	}
}
