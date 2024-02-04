package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigAddTelegramCmd struct{}

type ActorConfigAddTelegramRes struct{}

type ActorConfigAddTelegramHandler cqrs.HandlerFunc[ActorConfigAddTelegramCmd, *ActorConfigAddTelegramRes]

func NewActorConfigAddTelegramHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigAddTelegramHandler {
	return func(ctx context.Context, cmd ActorConfigAddTelegramCmd) (*ActorConfigAddTelegramRes, *i18np.Error) {
		return nil, nil
	}
}
