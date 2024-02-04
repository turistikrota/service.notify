package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigFilterQuery struct{}

type ActorConfigFilterRes struct{}

type ActorConfigFilterHandler cqrs.HandlerFunc[ActorConfigFilterQuery, *ActorConfigFilterRes]

func NewActorConfigFilterHandler(repo actor_config.Repository) ActorConfigFilterHandler {
	return func(ctx context.Context, cmd ActorConfigFilterQuery) (*ActorConfigFilterRes, *i18np.Error) {
		return nil, nil
	}
}
