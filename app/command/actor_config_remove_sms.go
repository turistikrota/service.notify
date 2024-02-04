package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigRemoveSmsCmd struct{}

type ActorConfigRemoveSmsRes struct{}

type ActorConfigRemoveSmsHandler cqrs.HandlerFunc[ActorConfigRemoveSmsCmd, *ActorConfigRemoveSmsRes]

func NewActorConfigRemoveSmsHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigRemoveSmsHandler {
	return func(ctx context.Context, cmd ActorConfigRemoveSmsCmd) (*ActorConfigRemoveSmsRes, *i18np.Error) {
		return nil, nil
	}
}
