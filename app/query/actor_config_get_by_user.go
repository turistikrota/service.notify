package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigGetByUserQuery struct {
	UserUUID string `params:"-" json:"-"`
	UserName string `params:"-" json:"-"`
}

type ActorConfigGetByUserRes struct {
	Detail *actor_config.UserDetailDto
}

type ActorConfigGetByUserHandler cqrs.HandlerFunc[ActorConfigGetByUserQuery, *ActorConfigGetByUserRes]

func NewActorConfigGetByUserHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigGetByUserHandler {
	return func(ctx context.Context, cmd ActorConfigGetByUserQuery) (*ActorConfigGetByUserRes, *i18np.Error) {
		res, err := repo.GetByUser(ctx, actor_config.WithActor{
			UUID: cmd.UserUUID,
			Name: cmd.UserName,
		})
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigGetByUserRes{
			Detail: res.ToUserDetail(),
		}, nil
	}
}
