package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigGetByBusinessUUIDQuery struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type ActorConfigGetByBusinessUUIDRes struct {
	Detail *actor_config.AdminDetailDto
}

type ActorConfigGetByBusinessUUIDHandler cqrs.HandlerFunc[ActorConfigGetByBusinessUUIDQuery, *ActorConfigGetByBusinessUUIDRes]

func NewActorConfigGetByBusinessUUIDHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigGetByBusinessUUIDHandler {
	return func(ctx context.Context, cmd ActorConfigGetByBusinessUUIDQuery) (*ActorConfigGetByBusinessUUIDRes, *i18np.Error) {
		res, err := repo.GetByBusinessUUID(ctx, cmd.UUID)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigGetByBusinessUUIDRes{
			Detail: res.ToAdminDetail(),
		}, nil
	}
}
