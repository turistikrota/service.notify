package event_stream

import (
	"context"

	"github.com/goccy/go-json"
	"github.com/turistikrota/service.notify/app/command"
)

func (h srv) OnAccountCreated(data []byte) {
	cmd := command.ActorConfigCreateUserCmd{}
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		return
	}
	h.app.Commands.ActorConfigCreateUser(context.TODO(), cmd)
}

func (h srv) OnBusinessCreated(data []byte) {
	cmd := command.ActorConfigCreateBusinessCmd{}
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		return
	}
	h.app.Commands.ActorConfigCreateBusiness(context.TODO(), cmd)
}

func (h srv) SendEmailToActor(data []byte) {
	cmd := command.NotifySendEmailCmd{}
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		return
	}
	h.app.Commands.NotifySendEmail(context.TODO(), cmd)
}

func (h srv) SendSmsToActor(data []byte) {

	cmd := command.NotifySendSmsCmd{}
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		return
	}
	h.app.Commands.NotifySendSms(context.TODO(), cmd)
}

func (h srv) SendSpecialEmail(data []byte) {
	cmd := command.NotifySendSpecialEmailCmd{}
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		return
	}
	h.app.Commands.NotifySendSpecialEmail(context.TODO(), cmd)
}

func (h srv) SendSpecialSms(data []byte) {
	cmd := command.NotifySendSpecialSmsCmd{}
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		return
	}
	h.app.Commands.NotifySendSpecialSms(context.TODO(), cmd)
}

func (h srv) SendNotification(data []byte) {
	cmd := command.NotifySendToAllChannelsCmd{}
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		return
	}
	h.app.Commands.NotifySendToAllChannels(context.TODO(), cmd)
}

func (h srv) SendPush(data []byte) {
	cmd := command.NotifySendPushCmd{}
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		return
	}
	h.app.Commands.NotifySendPush(context.TODO(), cmd)
}
