package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigRemoveMailCmd struct{}

type ActorConfigRemoveMailRes struct{}

type ActorConfigRemoveMailHandler cqrs.HandlerFunc[ActorConfigRemoveMailCmd, *ActorConfigRemoveMailRes]

func NewActorConfigRemoveMailHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigRemoveMailHandler {
	return func(ctx context.Context, cmd ActorConfigRemoveMailCmd) (*ActorConfigRemoveMailRes, *i18np.Error) {
		return nil, nil
	}
}
