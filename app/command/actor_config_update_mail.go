package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigUpdateMailCmd struct{}

type ActorConfigUpdateMailRes struct{}

type ActorConfigUpdateMailHandler cqrs.HandlerFunc[ActorConfigUpdateMailCmd, *ActorConfigUpdateMailRes]

func NewActorConfigUpdateMailHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigUpdateMailHandler {
	return func(ctx context.Context, cmd ActorConfigUpdateMailCmd) (*ActorConfigUpdateMailRes, *i18np.Error) {
		return nil, nil
	}
}
