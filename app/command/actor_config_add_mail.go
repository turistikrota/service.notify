package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigAddMailCmd struct{}

type ActorConfigAddMailRes struct{}

type ActorConfigAddMailHandler cqrs.HandlerFunc[ActorConfigAddMailCmd, *ActorConfigAddMailRes]

func NewActorConfigAddMailHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigAddMailHandler {
	return func(ctx context.Context, cmd ActorConfigAddMailCmd) (*ActorConfigAddMailRes, *i18np.Error) {
		return nil, nil
	}
}
