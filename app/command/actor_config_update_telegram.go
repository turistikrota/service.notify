package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigUpdateTelegramCmd struct {
	ActorUUID      string                           `json:"-"`
	ActorName      string                           `json:"-"`
	ActorType      actor_config.ActorType           `json:"-"`
	Credential     *actor_config.TelegramCredential `json:"credential" validate:"required,dive"`
	CredentialName string                           `json:"credential_name" validate:"required,min=3,max=100"`
}

type ActorConfigUpdateTelegramRes struct{}

type ActorConfigUpdateTelegramHandler cqrs.HandlerFunc[ActorConfigUpdateTelegramCmd, *ActorConfigUpdateTelegramRes]

func NewActorConfigUpdateTelegramHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigUpdateTelegramHandler {
	return func(ctx context.Context, cmd ActorConfigUpdateTelegramCmd) (*ActorConfigUpdateTelegramRes, *i18np.Error) {
		err := repo.UpdateTelegram(ctx, actor_config.Actor{
			UUID: cmd.ActorUUID,
			Name: cmd.ActorName,
			Type: cmd.ActorType,
		}, *cmd.Credential, cmd.CredentialName)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigUpdateTelegramRes{}, nil
	}
}
