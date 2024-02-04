package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigRemoveTelegramCmd struct{}

type ActorConfigRemoveTelegramRes struct{}

type ActorConfigRemoveTelegramHandler cqrs.HandlerFunc[ActorConfigRemoveTelegramCmd, *ActorConfigRemoveTelegramRes]

func NewActorConfigRemoveTelegramHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigRemoveTelegramHandler {
	return func(ctx context.Context, cmd ActorConfigRemoveTelegramCmd) (*ActorConfigRemoveTelegramRes, *i18np.Error) {
		return nil, nil
	}
}
