package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigAddMailCmd struct {
	ActorUUID  string                       `json:"-"`
	ActorName  string                       `json:"-"`
	ActorType  actor_config.ActorType       `json:"-"`
	Credential *actor_config.MailCredential `json:"credential" validate:"required,dive"`
}

type ActorConfigAddMailRes struct{}

type ActorConfigAddMailHandler cqrs.HandlerFunc[ActorConfigAddMailCmd, *ActorConfigAddMailRes]

func NewActorConfigAddMailHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigAddMailHandler {
	return func(ctx context.Context, cmd ActorConfigAddMailCmd) (*ActorConfigAddMailRes, *i18np.Error) {
		err := repo.AddMail(ctx, actor_config.Actor{
			UUID: cmd.ActorUUID,
			Name: cmd.ActorName,
			Type: cmd.ActorType,
		}, *cmd.Credential)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigAddMailRes{}, nil
	}
}
