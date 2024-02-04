package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigGetByUUIDQuery struct{}

type ActorConfigGetByUUIDRes struct{}

type ActorConfigGetByUUIDHandler cqrs.HandlerFunc[ActorConfigGetByUUIDQuery, *ActorConfigGetByUUIDRes]

func NewActorConfigGetByUUIDHandler(repo actor_config.Repository) ActorConfigGetByUUIDHandler {
	return func(ctx context.Context, cmd ActorConfigGetByUUIDQuery) (*ActorConfigGetByUUIDRes, *i18np.Error) {
		return nil, nil
	}
}
