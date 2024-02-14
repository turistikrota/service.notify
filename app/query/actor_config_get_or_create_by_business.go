package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigGetOrCreateByBusinessQuery struct {
	BusinessUUID string `params:"-" json:"-"`
	BusinessName string `params:"-" json:"-"`
}

type ActorConfigGetOrCreateByBusinessRes struct {
	Detail *actor_config.BusinessDetailDto
}

type ActorConfigGetOrCreateByBusinessHandler cqrs.HandlerFunc[ActorConfigGetOrCreateByBusinessQuery, *ActorConfigGetOrCreateByBusinessRes]

func NewActorConfigGetOrCreateByBusinessHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigGetOrCreateByBusinessHandler {
	return func(ctx context.Context, cmd ActorConfigGetOrCreateByBusinessQuery) (*ActorConfigGetOrCreateByBusinessRes, *i18np.Error) {
		res, err := repo.GetByBusinessOrCreate(ctx, actor_config.WithActor{
			UUID: cmd.BusinessUUID,
			Name: cmd.BusinessName,
		})
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigGetOrCreateByBusinessRes{
			Detail: res.ToBusinessDetail(),
		}, nil
	}
}
