package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/adapters/push"
	"github.com/turistikrota/service.notify/domains/notify"
	"github.com/turistikrota/service.shared/auth/session"
)

type NotifySendPushCmd struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	Image     string `json:"image"`
	UserUUID  string `json:"userUUID"`
	Translate bool   `json:"translate"`
}

type NotifySendPushRes struct{}

type NotifySendPushHandler cqrs.HandlerFunc[NotifySendPushCmd, *NotifySendPushRes]

func NewNotifySendPushHandler(factory notify.Factory, i18n *i18np.I18n, srv push.Service, sessionSrv session.Service) NotifySendPushHandler {
	return func(ctx context.Context, cmd NotifySendPushCmd) (*NotifySendPushRes, *i18np.Error) {
		sess, err := sessionSrv.GetAll(cmd.UserUUID)
		if err != nil {
			return nil, err
		}
		go func() {
			for _, s := range sess {
				if s.FcmToken != "" {
					err := srv.Send(ctx, push.SendConfig{
						Token: s.FcmToken,
						Title: cmd.Title,
						Body:  cmd.Body,
						Image: cmd.Image,
					})
					if err != nil {
						return
					}
				}
			}
		}()
		return &NotifySendPushRes{}, nil
	}
}
