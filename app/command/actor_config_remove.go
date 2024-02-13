package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigRemoveCmd struct {
	ActorUUID      string                 `json:"-"`
	ActorName      string                 `json:"-"`
	ActorType      actor_config.ActorType `json:"-"`
	Type           string                 `json:"type" validate:"required,oneof=mail sms telegram"`
	CredentialName string                 `json:"credential_name" validate:"required,min=3,max=100"`
}

type ActorConfigRemoveRes struct{}

type ActorConfigRemoveHandler cqrs.HandlerFunc[ActorConfigRemoveCmd, *ActorConfigRemoveRes]

func NewActorConfigRemoveHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigRemoveHandler {
	return func(ctx context.Context, cmd ActorConfigRemoveCmd) (*ActorConfigRemoveRes, *i18np.Error) {
		var err *i18np.Error
		if cmd.Type == actor_config.TypeMail.String() {
			err = repo.RemoveMail(ctx, actor_config.Actor{
				UUID: cmd.ActorUUID,
				Name: cmd.ActorName,
				Type: cmd.ActorType,
			}, cmd.CredentialName)
		}
		if cmd.Type == actor_config.TypeSMS.String() {
			err = repo.RemoveSMS(ctx, actor_config.Actor{
				UUID: cmd.ActorUUID,
				Name: cmd.ActorName,
				Type: cmd.ActorType,
			}, cmd.CredentialName)
		}
		if cmd.Type == actor_config.TypeTelegram.String() {
			err = repo.RemoveTelegram(ctx, actor_config.Actor{
				UUID: cmd.ActorUUID,
				Name: cmd.ActorName,
				Type: cmd.ActorType,
			}, cmd.CredentialName)
		}
		if err != nil {
			return nil, err
		}
		return &ActorConfigRemoveRes{}, nil
	}
}
