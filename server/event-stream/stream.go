package event_stream

import (
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/server"
	"github.com/turistikrota/service.notify/app"
	"github.com/turistikrota/service.notify/config"
)

type srv struct {
	app    app.Application
	topics config.Topics
	engine events.Engine
}

type Config struct {
	App    app.Application
	Engine events.Engine
	Topics config.Topics
}

func New(config Config) server.Server {
	return srv{
		app:    config.App,
		engine: config.Engine,
		topics: config.Topics,
	}
}

func (s srv) Listen() error {
	s.engine.Subscribe(s.topics.Account.Created, s.OnAccountCreated)
	s.engine.Subscribe(s.topics.Business.Created, s.OnBusinessCreated)
	s.engine.Subscribe(s.topics.Notify.SendEmailToActor, s.SendEmailToActor)
	s.engine.Subscribe(s.topics.Notify.SendSmsToActor, s.SendSmsToActor)
	s.engine.Subscribe(s.topics.Notify.SendSpecialEmail, s.SendSpecialEmail)
	s.engine.Subscribe(s.topics.Notify.SendSpecialSms, s.SendSpecialSms)
	s.engine.Subscribe(s.topics.Notify.SendNotification, s.SendNotification)
	s.engine.Subscribe(s.topics.Notify.SendPush, s.SendPush)
	return nil
}
