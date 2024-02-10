package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigAddSmsCmd struct {
	ActorUUID  string                      `json:"-"`
	ActorName  string                      `json:"-"`
	ActorType  actor_config.ActorType      `json:"-"`
	Credential *actor_config.SMSCredential `json:"credential" validate:"required,dive"`
}

type ActorConfigAddSmsRes struct{}

type ActorConfigAddSmsHandler cqrs.HandlerFunc[ActorConfigAddSmsCmd, *ActorConfigAddSmsRes]

func NewActorConfigAddSmsHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigAddSmsHandler {
	return func(ctx context.Context, cmd ActorConfigAddSmsCmd) (*ActorConfigAddSmsRes, *i18np.Error) {
		err := repo.AddSMS(ctx, actor_config.Actor{
			UUID: cmd.ActorUUID,
			Name: cmd.ActorName,
			Type: cmd.ActorType,
		}, *cmd.Credential)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigAddSmsRes{}, nil
	}
}
