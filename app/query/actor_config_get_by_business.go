package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigGetByBusinessQuery struct {
	BusinessUUID string `params:"-" json:"-"`
	BusinessName string `params:"-" json:"-"`
}

type ActorConfigGetByBusinessRes struct {
	Detail *actor_config.BusinessDetailDto
}

type ActorConfigGetByBusinessHandler cqrs.HandlerFunc[ActorConfigGetByBusinessQuery, *ActorConfigGetByBusinessRes]

func NewActorConfigGetByBusinessHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigGetByBusinessHandler {
	return func(ctx context.Context, cmd ActorConfigGetByBusinessQuery) (*ActorConfigGetByBusinessRes, *i18np.Error) {
		res, err := repo.GetByBusiness(ctx, actor_config.WithActor{
			UUID: cmd.BusinessUUID,
			Name: cmd.BusinessName,
		})
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigGetByBusinessRes{
			Detail: res.ToBusinessDetail(),
		}, nil
	}
}
