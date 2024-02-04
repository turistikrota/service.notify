package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigGetByUserQuery struct{}

type ActorConfigGetByUserRes struct{}

type ActorConfigGetByUserHandler cqrs.HandlerFunc[ActorConfigGetByUserQuery, *ActorConfigGetByUserRes]

func NewActorConfigGetByUserHandler(repo actor_config.Repository) ActorConfigGetByUserHandler {
	return func(ctx context.Context, cmd ActorConfigGetByUserQuery) (*ActorConfigGetByUserRes, *i18np.Error) {
		return nil, nil
	}
}
