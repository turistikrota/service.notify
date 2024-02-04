package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigGetByBusinessQuery struct{}

type ActorConfigGetByBusinessRes struct{}

type ActorConfigGetByBusinessHandler cqrs.HandlerFunc[ActorConfigGetByBusinessQuery, *ActorConfigGetByBusinessRes]

func NewActorConfigGetByBusinessHandler(repo actor_config.Repository) ActorConfigGetByBusinessHandler {
	return func(ctx context.Context, cmd ActorConfigGetByBusinessQuery) (*ActorConfigGetByBusinessRes, *i18np.Error) {
		return nil, nil
	}
}
