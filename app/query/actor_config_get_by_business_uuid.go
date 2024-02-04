package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigGetByBusinessUUIDQuery struct{}

type ActorConfigGetByBusinessUUIDRes struct{}

type ActorConfigGetByBusinessUUIDHandler cqrs.HandlerFunc[ActorConfigGetByBusinessUUIDQuery, *ActorConfigGetByBusinessUUIDRes]

func NewActorConfigGetByBusinessUUIDHandler(repo actor_config.Repository) ActorConfigGetByBusinessUUIDHandler {
	return func(ctx context.Context, cmd ActorConfigGetByBusinessUUIDQuery) (*ActorConfigGetByBusinessUUIDRes, *i18np.Error) {
		return nil, nil
	}
}
