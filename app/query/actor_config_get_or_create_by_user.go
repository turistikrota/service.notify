package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigGetOrCreateByUserQuery struct {
	UserUUID string `params:"-" json:"-"`
	UserName string `params:"-" json:"-"`
}

type ActorConfigGetOrCreateByUserRes struct {
	Detail *actor_config.UserDetailDto
}

type ActorConfigGetOrCreateByUserHandler cqrs.HandlerFunc[ActorConfigGetOrCreateByUserQuery, *ActorConfigGetOrCreateByUserRes]

func NewActorConfigGetOrCreateByUserHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigGetOrCreateByUserHandler {
	return func(ctx context.Context, cmd ActorConfigGetOrCreateByUserQuery) (*ActorConfigGetOrCreateByUserRes, *i18np.Error) {
		res, err := repo.GetByUserOrCreate(ctx, actor_config.WithActor{
			UUID: cmd.UserUUID,
			Name: cmd.UserName,
		})
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigGetOrCreateByUserRes{
			Detail: res.ToUserDetail(),
		}, nil
	}
}
