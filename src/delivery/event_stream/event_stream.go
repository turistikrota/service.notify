package event_stream

import (
	"context"

	"api.turistikrota.com/notify/src/app"
	"api.turistikrota.com/notify/src/config"
	"api.turistikrota.com/notify/src/delivery/event_stream/dto"
	"github.com/turistikrota/service.shared/events"
	"github.com/sirupsen/logrus"
)

type Server struct {
	app    app.Application
	Topics config.NotifyTopics
	engine events.Engine
	ctx    context.Context
	dto    dto.Dto
}

type Config struct {
	App    app.Application
	Topics config.NotifyTopics
	Engine events.Engine
	Ctx    context.Context
}

func New(config Config) Server {
	return Server{
		app:    config.App,
		engine: config.Engine,
		Topics: config.Topics,
		ctx:    config.Ctx,
		dto:    dto.New(),
	}
}

func (s Server) Load() {
	logrus.Info("Loading event stream server")
	s.engine.Subscribe(s.Topics.Email, s.ListenSendNotifyEmail)
	s.engine.Subscribe(s.Topics.SMS, s.ListenSendNotifySMS)
	s.engine.Subscribe(s.Topics.Telegram, s.ListenSendNotifyTelegram)
}
