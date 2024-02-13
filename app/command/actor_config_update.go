package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/domains/actor_config"
)

type ActorConfigUpdateCmd struct {
	ActorUUID      string                 `json:"-"`
	ActorName      string                 `json:"-"`
	ActorType      actor_config.ActorType `json:"-"`
	Name           string                 `json:"name" validate:"required,min=3,max=100"`
	Type           string                 `json:"type" validate:"required,oneof=mail sms telegram"`
	ChatID         *string                `json:"chatId" validate:"required_if=Type telegram,omitempty,min=3,max=100"`
	Email          *string                `json:"email" validate:"required_if=Type mail,omitempty,email"`
	Phone          *string                `json:"phone" validate:"required_if=Type sms,omitempty,e164"`
	CredentialName string                 `json:"credential_name" validate:"required,min=3,max=100"`
}

type ActorConfigUpdateRes struct{}

type ActorConfigUpdateHandler cqrs.HandlerFunc[ActorConfigUpdateCmd, *ActorConfigUpdateRes]

func NewActorConfigUpdateHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigUpdateHandler {
	return func(ctx context.Context, cmd ActorConfigUpdateCmd) (*ActorConfigUpdateRes, *i18np.Error) {
		var err *i18np.Error
		if cmd.Type == actor_config.TypeMail.String() {
			err = repo.UpdateMail(ctx, actor_config.Actor{
				UUID: cmd.ActorUUID,
				Name: cmd.ActorName,
				Type: cmd.ActorType,
			}, actor_config.MailCredential{
				Name:  cmd.Name,
				Email: *cmd.Email,
			}, cmd.CredentialName)
		}
		if cmd.Type == actor_config.TypeSMS.String() {
			p := *cmd.Phone
			countryCode := p[:3]
			phone := p[3:]
			err = repo.UpdateSMS(ctx, actor_config.Actor{
				UUID: cmd.ActorUUID,
				Name: cmd.ActorName,
				Type: cmd.ActorType,
			}, actor_config.SMSCredential{
				Name:        cmd.Name,
				Phone:       phone,
				CountryCode: countryCode,
			}, cmd.CredentialName)
		}
		if cmd.Type == actor_config.TypeTelegram.String() {
			err = repo.UpdateTelegram(ctx, actor_config.Actor{
				UUID: cmd.ActorUUID,
				Name: cmd.ActorName,
				Type: cmd.ActorType,
			}, actor_config.TelegramCredential{
				Name:   cmd.Name,
				ChatID: *cmd.ChatID,
			}, cmd.CredentialName)
		}
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ActorConfigUpdateRes{}, nil
	}
}
