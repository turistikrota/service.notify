package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigCreateCmd struct{}

type ActorConfigCreateRes struct{}

type ActorConfigCreateHandler cqrs.HandlerFunc[ActorConfigCreateCmd, *ActorConfigCreateRes]

func NewActorConfigCreateHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigCreateHandler {
	return func(ctx context.Context, cmd ActorConfigCreateCmd) (*ActorConfigCreateRes, *i18np.Error) {
		return nil, nil
	}
}
