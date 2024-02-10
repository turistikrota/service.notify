package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigRemoveTelegramCmd struct {
	ActorUUID      string                 `json:"-"`
	ActorName      string                 `json:"-"`
	ActorType      actor_config.ActorType `json:"-"`
	CredentialName string                 `json:"credential_name" validate:"required,min=3,max=100"`
}

type ActorConfigRemoveTelegramRes struct{}

type ActorConfigRemoveTelegramHandler cqrs.HandlerFunc[ActorConfigRemoveTelegramCmd, *ActorConfigRemoveTelegramRes]

func NewActorConfigRemoveTelegramHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigRemoveTelegramHandler {
	return func(ctx context.Context, cmd ActorConfigRemoveTelegramCmd) (*ActorConfigRemoveTelegramRes, *i18np.Error) {
		err := repo.RemoveTelegram(ctx, actor_config.Actor{
			UUID: cmd.ActorUUID,
			Name: cmd.ActorName,
			Type: cmd.ActorType,
		}, cmd.CredentialName)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigRemoveTelegramRes{}, nil
	}
}
