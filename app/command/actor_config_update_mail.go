package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigUpdateMailCmd struct {
	ActorUUID      string                       `json:"-"`
	ActorName      string                       `json:"-"`
	ActorType      actor_config.ActorType       `json:"-"`
	Credential     *actor_config.MailCredential `json:"credential" validate:"required,dive"`
	CredentialName string                       `json:"credential_name" validate:"required,min=3,max=100"`
}

type ActorConfigUpdateMailRes struct{}

type ActorConfigUpdateMailHandler cqrs.HandlerFunc[ActorConfigUpdateMailCmd, *ActorConfigUpdateMailRes]

func NewActorConfigUpdateMailHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigUpdateMailHandler {
	return func(ctx context.Context, cmd ActorConfigUpdateMailCmd) (*ActorConfigUpdateMailRes, *i18np.Error) {
		err := repo.UpdateMail(ctx, actor_config.Actor{
			UUID: cmd.ActorUUID,
			Name: cmd.ActorName,
			Type: cmd.ActorType,
		}, *cmd.Credential, cmd.CredentialName)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigUpdateMailRes{}, nil
	}
}
