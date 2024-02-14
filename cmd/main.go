package main

import (
	"github.com/cilloparch/cillop/db/mredis"
	"github.com/cilloparch/cillop/env"
	"github.com/cilloparch/cillop/events/nats"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/validation"
	"github.com/turistikrota/service.notify/config"
	event_stream "github.com/turistikrota/service.notify/server/event-stream"
	"github.com/turistikrota/service.notify/server/http"
	"github.com/turistikrota/service.notify/service"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/db/mongo"
	"github.com/turistikrota/service.shared/db/redis"
)

func main() {
	cnf := config.App{}
	env.Load(&cnf)
	i18n := i18np.New(cnf.I18n.Fallback)
	i18n.Load(cnf.I18n.Dir, cnf.I18n.Locales...)
	eventEngine := nats.New(nats.Config{
		Url:     cnf.Nats.Url,
		Streams: cnf.Nats.Streams,
	})
	valid := validation.New(i18n)
	valid.ConnectCustom()
	valid.RegisterTagName()
	mongo := loadMongo(cnf)
	cache := mredis.New(&mredis.Config{
		Host:     cnf.CacheRedis.Host,
		Port:     cnf.CacheRedis.Port,
		Password: cnf.CacheRedis.Pw,
		DB:       cnf.CacheRedis.Db,
	})
	r := redis.New(&redis.Config{
		Host:     cnf.Redis.Host,
		Port:     cnf.Redis.Port,
		Password: cnf.Redis.Pw,
		DB:       cnf.Redis.Db,
	})
	tknSrv := token.New(token.Config{
		Expiration:     cnf.TokenSrv.Expiration,
		PublicKeyFile:  cnf.RSA.PublicKeyFile,
		PrivateKeyFile: cnf.RSA.PrivateKeyFile,
	})
	session := session.NewSessionApp(session.Config{
		Redis:       r,
		EventEngine: eventEngine,
		TokenSrv:    tknSrv,
		Topic:       cnf.Session.Topic,
		Project:     cnf.TokenSrv.Project,
	})
	app := service.NewApplication(service.Config{
		App:         cnf,
		EventEngine: eventEngine,
		CacheSrv:    cache,
		Validator:   valid,
		MongoDB:     mongo,
		I18n:        i18n,
		SessionSrv:  session.Service,
	})
	http := http.New(http.Config{
		Env:         cnf,
		App:         app,
		I18n:        i18n,
		Validator:   *valid,
		HttpHeaders: cnf.HttpHeaders,
		TokenSrv:    tknSrv,
		SessionSrv:  session.Service,
	})
	eventStream := event_stream.New(event_stream.Config{
		App:    app,
		Engine: eventEngine,
		Topics: cnf.Topics,
	})
	go eventStream.Listen()
	http.Listen()
}

func loadMongo(cnf config.App) *mongo.DB {
	uri := mongo.CalcMongoUri(mongo.UriParams{
		Host:  cnf.DB.ActorConfig.Host,
		Port:  cnf.DB.ActorConfig.Port,
		User:  cnf.DB.ActorConfig.Username,
		Pass:  cnf.DB.ActorConfig.Password,
		Db:    cnf.DB.ActorConfig.Database,
		Query: cnf.DB.ActorConfig.Query,
	})
	d, err := mongo.New(uri, cnf.DB.ActorConfig.Database)
	if err != nil {
		panic(err)
	}
	return d
}
