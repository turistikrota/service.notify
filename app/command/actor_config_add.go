package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigAddCmd struct {
	ActorUUID string                 `json:"-"`
	ActorName string                 `json:"-"`
	ActorType actor_config.ActorType `json:"-"`
	Name      string                 `json:"name" validate:"required,min=3,max=100"`
	Type      string                 `json:"type" validate:"required,oneof=mail sms telegram"`
	ChatID    *string                `json:"chatId" validate:"required_if=Type telegram,omitempty,min=3,max=100"`
	Email     *string                `json:"email" validate:"required_if=Type mail,omitempty,email"`
	Phone     *string                `json:"phone" validate:"required_if=Type sms,omitempty,e164"`
}

type ActorConfigAddRes struct{}

type ActorConfigAddHandler cqrs.HandlerFunc[ActorConfigAddCmd, *ActorConfigAddRes]

func NewActorConfigAddHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigAddHandler {
	return func(ctx context.Context, cmd ActorConfigAddCmd) (*ActorConfigAddRes, *i18np.Error) {
		var err *i18np.Error
		if cmd.Type == actor_config.TypeMail.String() {
			err = repo.AddMail(ctx, actor_config.Actor{
				UUID: cmd.ActorUUID,
				Name: cmd.ActorName,
				Type: cmd.ActorType,
			}, actor_config.MailCredential{
				Name:  cmd.Name,
				Email: *cmd.Email,
			})
		}
		if cmd.Type == actor_config.TypeSMS.String() {
			p := *cmd.Phone
			countryCode := p[:3]
			phone := p[3:]
			err = repo.AddSMS(ctx, actor_config.Actor{
				UUID: cmd.ActorUUID,
				Name: cmd.ActorName,
				Type: cmd.ActorType,
			}, actor_config.SMSCredential{
				Name:        cmd.Name,
				Phone:       phone,
				CountryCode: countryCode,
			})
		}
		if cmd.Type == actor_config.TypeTelegram.String() {
			err = repo.AddTelegram(ctx, actor_config.Actor{
				UUID: cmd.ActorUUID,
				Name: cmd.ActorName,
				Type: cmd.ActorType,
			}, actor_config.TelegramCredential{
				Name:   cmd.Name,
				ChatID: *cmd.ChatID,
			})
		}
		if err != nil {
			return nil, err
		}
		return &ActorConfigAddRes{}, nil
	}
}
