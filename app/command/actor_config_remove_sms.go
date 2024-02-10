package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigRemoveSmsCmd struct {
	ActorUUID      string                 `json:"-"`
	ActorName      string                 `json:"-"`
	ActorType      actor_config.ActorType `json:"-"`
	CredentialName string                 `json:"credential_name" validate:"required,min=3,max=100"`
}

type ActorConfigRemoveSmsRes struct{}

type ActorConfigRemoveSmsHandler cqrs.HandlerFunc[ActorConfigRemoveSmsCmd, *ActorConfigRemoveSmsRes]

func NewActorConfigRemoveSmsHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigRemoveSmsHandler {
	return func(ctx context.Context, cmd ActorConfigRemoveSmsCmd) (*ActorConfigRemoveSmsRes, *i18np.Error) {
		err := repo.RemoveSMS(ctx, actor_config.Actor{
			UUID: cmd.ActorUUID,
			Name: cmd.ActorName,
			Type: cmd.ActorType,
		}, cmd.CredentialName)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigRemoveSmsRes{}, nil
	}
}
