package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigUpdateSmsCmd struct {
	ActorUUID      string                      `json:"-"`
	ActorName      string                      `json:"-"`
	ActorType      actor_config.ActorType      `json:"-"`
	Credential     *actor_config.SMSCredential `json:"credential" validate:"required,dive"`
	CredentialName string                      `json:"credential_name" validate:"required,min=3,max=100"`
}

type ActorConfigUpdateSmsRes struct{}

type ActorConfigUpdateSmsHandler cqrs.HandlerFunc[ActorConfigUpdateSmsCmd, *ActorConfigUpdateSmsRes]

func NewActorConfigUpdateSmsHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigUpdateSmsHandler {
	return func(ctx context.Context, cmd ActorConfigUpdateSmsCmd) (*ActorConfigUpdateSmsRes, *i18np.Error) {
		err := repo.UpdateSMS(ctx, actor_config.Actor{
			UUID: cmd.ActorUUID,
			Name: cmd.ActorName,
			Type: cmd.ActorType,
		}, *cmd.Credential, cmd.CredentialName)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigUpdateSmsRes{}, nil
	}
}
