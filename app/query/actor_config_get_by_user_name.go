package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigGetByUserNameQuery struct {
	Name string `params:"name" validate:"required"`
}

type ActorConfigGetByUserNameRes struct {
	Detail *actor_config.AdminDetailDto
}

type ActorConfigGetByUserNameHandler cqrs.HandlerFunc[ActorConfigGetByUserNameQuery, *ActorConfigGetByUserNameRes]

func NewActorConfigGetByUserNameHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigGetByUserNameHandler {
	return func(ctx context.Context, cmd ActorConfigGetByUserNameQuery) (*ActorConfigGetByUserNameRes, *i18np.Error) {
		res, err := repo.GetByUserName(ctx, cmd.Name)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigGetByUserNameRes{
			Detail: res.ToAdminDetail(),
		}, nil
	}
}
