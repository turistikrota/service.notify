package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigGetByUserUUIDQuery struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type ActorConfigGetByUserUUIDRes struct {
	Detail *actor_config.AdminDetailDto
}

type ActorConfigGetByUserUUIDHandler cqrs.HandlerFunc[ActorConfigGetByUserUUIDQuery, *ActorConfigGetByUserUUIDRes]

func NewActorConfigGetByUserUUIDHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigGetByUserUUIDHandler {
	return func(ctx context.Context, cmd ActorConfigGetByUserUUIDQuery) (*ActorConfigGetByUserUUIDRes, *i18np.Error) {
		res, err := repo.GetByUserUUID(ctx, cmd.UUID)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigGetByUserUUIDRes{
			Detail: res.ToAdminDetail(),
		}, nil
	}
}
