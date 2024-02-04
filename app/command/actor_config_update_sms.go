package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigUpdateSmsCmd struct{}

type ActorConfigUpdateSmsRes struct{}

type ActorConfigUpdateSmsHandler cqrs.HandlerFunc[ActorConfigUpdateSmsCmd, *ActorConfigUpdateSmsRes]

func NewActorConfigUpdateSmsHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigUpdateSmsHandler {
	return func(ctx context.Context, cmd ActorConfigUpdateSmsCmd) (*ActorConfigUpdateSmsRes, *i18np.Error) {
		return nil, nil
	}
}
