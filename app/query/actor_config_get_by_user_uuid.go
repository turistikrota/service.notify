package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigGetByUserUUIDQuery struct{}

type ActorConfigGetByUserUUIDRes struct{}

type ActorConfigGetByUserUUIDHandler cqrs.HandlerFunc[ActorConfigGetByUserUUIDQuery, *ActorConfigGetByUserUUIDRes]

func NewActorConfigGetByUserUUIDHandler(repo actor_config.Repository) ActorConfigGetByUserUUIDHandler {
	return func(ctx context.Context, cmd ActorConfigGetByUserUUIDQuery) (*ActorConfigGetByUserUUIDRes, *i18np.Error) {
		return nil, nil
	}
}
