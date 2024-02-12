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
