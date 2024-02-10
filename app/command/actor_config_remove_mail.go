package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigRemoveMailCmd struct {
	ActorUUID      string                 `json:"-"`
	ActorName      string                 `json:"-"`
	ActorType      actor_config.ActorType `json:"-"`
	CredentialName string                 `json:"credential_name" validate:"required,min=3,max=100"`
}

type ActorConfigRemoveMailRes struct{}

type ActorConfigRemoveMailHandler cqrs.HandlerFunc[ActorConfigRemoveMailCmd, *ActorConfigRemoveMailRes]

func NewActorConfigRemoveMailHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigRemoveMailHandler {
	return func(ctx context.Context, cmd ActorConfigRemoveMailCmd) (*ActorConfigRemoveMailRes, *i18np.Error) {
		err := repo.RemoveMail(ctx, actor_config.Actor{
			UUID: cmd.ActorUUID,
			Name: cmd.ActorName,
			Type: cmd.ActorType,
		}, cmd.CredentialName)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigRemoveMailRes{}, nil
	}
}
