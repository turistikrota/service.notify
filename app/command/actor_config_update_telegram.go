package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigUpdateTelegramCmd struct{}

type ActorConfigUpdateTelegramRes struct{}

type ActorConfigUpdateTelegramHandler cqrs.HandlerFunc[ActorConfigUpdateTelegramCmd, *ActorConfigUpdateTelegramRes]

func NewActorConfigUpdateTelegramHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigUpdateTelegramHandler {
	return func(ctx context.Context, cmd ActorConfigUpdateTelegramCmd) (*ActorConfigUpdateTelegramRes, *i18np.Error) {
		return nil, nil
	}
}
