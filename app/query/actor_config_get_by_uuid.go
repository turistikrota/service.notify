package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigGetByUUIDQuery struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type ActorConfigGetByUUIDRes struct {
	Detail *actor_config.AdminDetailDto
}

type ActorConfigGetByUUIDHandler cqrs.HandlerFunc[ActorConfigGetByUUIDQuery, *ActorConfigGetByUUIDRes]

func NewActorConfigGetByUUIDHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigGetByUUIDHandler {
	return func(ctx context.Context, cmd ActorConfigGetByUUIDQuery) (*ActorConfigGetByUUIDRes, *i18np.Error) {
		res, err := repo.GetByUUID(ctx, cmd.UUID)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigGetByUUIDRes{
			Detail: res.ToAdminDetail(),
		}, nil
	}
}
