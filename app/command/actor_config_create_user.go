package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigCreateUserCmd struct {
	UserUUID string `json:"user_uuid"`
	UserName string `json:"name"`
}

type ActorConfigCreateUserRes struct{}

type ActorConfigCreateUserHandler cqrs.HandlerFunc[ActorConfigCreateUserCmd, *ActorConfigCreateUserRes]

func NewActorConfigCreateUserHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigCreateUserHandler {
	return func(ctx context.Context, cmd ActorConfigCreateUserCmd) (*ActorConfigCreateUserRes, *i18np.Error) {
		err := repo.Create(ctx, factory.New(actor_config.Actor{
			UUID: cmd.UserUUID,
			Name: cmd.UserName,
			Type: actor_config.ActorTypeUser,
		}))
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigCreateUserRes{}, nil
	}
}
