package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigCreateBusinessCmd struct {
	UUID     string `json:"uuid"`
	NickName string `json:"nickName"`
}

type ActorConfigCreateBusinessRes struct{}

type ActorConfigCreateBusinessHandler cqrs.HandlerFunc[ActorConfigCreateBusinessCmd, *ActorConfigCreateBusinessRes]

func NewActorConfigCreateBusinessHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigCreateBusinessHandler {
	return func(ctx context.Context, cmd ActorConfigCreateBusinessCmd) (*ActorConfigCreateBusinessRes, *i18np.Error) {
		err := repo.Create(ctx, factory.New(actor_config.Actor{
			UUID: cmd.UUID,
			Name: cmd.NickName,
			Type: actor_config.ActorTypeBusiness,
		}))
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigCreateBusinessRes{}, nil
	}
}
