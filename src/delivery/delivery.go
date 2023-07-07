package delivery

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.notify/src/app"
	"github.com/turistikrota/service.notify/src/config"
	"github.com/turistikrota/service.notify/src/delivery/event_stream"
	"github.com/turistikrota/service.notify/src/delivery/http"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/events"
	sharedHttp "github.com/turistikrota/service.shared/server/http"
	"github.com/turistikrota/service.shared/validator"
)

type Delivery interface {
	Load()
}

type delivery struct {
	app         app.Application
	config      config.App
	i18n        *i18np.I18n
	ctx         context.Context
	eventEngine events.Engine
	tknSrv      token.Service
	sessionSrv  session.Service
}

type Config struct {
	App         app.Application
	Config      config.App
	I18n        *i18np.I18n
	Ctx         context.Context
	EventEngine events.Engine
	TokenSrv    token.Service
	SessionSrv  session.Service
}

func New(config Config) Delivery {
	return &delivery{
		app:         config.App,
		config:      config.Config,
		i18n:        config.I18n,
		ctx:         config.Ctx,
		eventEngine: config.EventEngine,
		tknSrv:      config.TokenSrv,
		sessionSrv:  config.SessionSrv,
	}
}

func (d *delivery) Load() {
	d.loadEventStream().loadHTTP()
}

func (d *delivery) loadHTTP() *delivery {
	sharedHttp.RunServer(sharedHttp.Config{
		Host:  d.config.Server.Host,
		Port:  d.config.Server.Port,
		I18n:  d.i18n,
		Group: d.config.Server.Group,
		CreateHandler: func(router fiber.Router) fiber.Router {
			val := validator.New(d.i18n)
			val.ConnectCustom()
			val.RegisterTagName()
			return http.New(http.Config{
				App:         d.app,
				I18n:        *d.i18n,
				Validator:   *val,
				Context:     d.ctx,
				HttpHeaders: d.config.HttpHeaders,
				TokenSrv:    d.tknSrv,
				SessionSrv:  d.sessionSrv,
			}).Load(router)
		},
	})
	return d
}

func (d *delivery) loadEventStream() *delivery {
	eventStream := event_stream.New(event_stream.Config{
		App:    d.app,
		Topics: d.config.Topics.Notify,
		Engine: d.eventEngine,
		Ctx:    d.ctx,
	})
	err := d.eventEngine.Open()
	if err != nil {
		panic(err)
	}
	eventStream.Load()
	return d
}
