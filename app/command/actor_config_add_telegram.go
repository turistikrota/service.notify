package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigAddTelegramCmd struct {
	ActorUUID  string                           `json:"-"`
	ActorName  string                           `json:"-"`
	ActorType  actor_config.ActorType           `json:"-"`
	Credential *actor_config.TelegramCredential `json:"credential" validate:"required,dive"`
}

type ActorConfigAddTelegramRes struct{}

type ActorConfigAddTelegramHandler cqrs.HandlerFunc[ActorConfigAddTelegramCmd, *ActorConfigAddTelegramRes]

func NewActorConfigAddTelegramHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigAddTelegramHandler {
	return func(ctx context.Context, cmd ActorConfigAddTelegramCmd) (*ActorConfigAddTelegramRes, *i18np.Error) {
		err := repo.AddTelegram(ctx, actor_config.Actor{
			UUID: cmd.ActorUUID,
			Name: cmd.ActorName,
			Type: cmd.ActorType,
		}, *cmd.Credential)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigAddTelegramRes{}, nil
	}
}
