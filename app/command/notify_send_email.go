package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
)

type NotifySendEmailCmd struct{}

type NotifySendEmailRes struct{}

type NotifySendEmailHandler cqrs.HandlerFunc[NotifySendEmailCmd, *NotifySendEmailRes]

func NewNotifySendEmailHandler() NotifySendEmailHandler {
	return func(ctx context.Context, cmd NotifySendEmailCmd) (*NotifySendEmailRes, *i18np.Error) {
		return &NotifySendEmailRes{}, nil
	}
}
