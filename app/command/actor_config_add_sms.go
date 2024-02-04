package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigAddSmsCmd struct{}

type ActorConfigAddSmsRes struct{}

type ActorConfigAddSmsHandler cqrs.HandlerFunc[ActorConfigAddSmsCmd, *ActorConfigAddSmsRes]

func NewActorConfigAddSmsHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigAddSmsHandler {
	return func(ctx context.Context, cmd ActorConfigAddSmsCmd) (*ActorConfigAddSmsRes, *i18np.Error) {
		return nil, nil
	}
}
