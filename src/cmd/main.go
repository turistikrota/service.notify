package main

import (
	"context"
	"fmt"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/env"
	"github.com/mixarchitecture/microp/events/nats"
	"github.com/mixarchitecture/microp/logs"
	"github.com/mixarchitecture/microp/validator"
	"github.com/turistikrota/service.notify/src/config"
	"github.com/turistikrota/service.notify/src/delivery"
	"github.com/turistikrota/service.notify/src/service"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/db/mongo"
	"github.com/turistikrota/service.shared/db/redis"
)

func main() {
	logs.Init()
	ctx := context.Background()
	config := config.App{}
	env.Load(&config)
	i18n := i18np.New(config.I18n.Fallback)
	i18n.Load(config.I18n.Dir, config.I18n.Locales...)
	eventEngine := nats.New(nats.Config{
		Url:     config.Nats.Url,
		Streams: config.Nats.Streams,
	})
	notifyMongo := loadNotifyMongo(config)
	valid := validator.New(i18n)
	valid.ConnectCustom()
	valid.RegisterTagName()
	redis := redis.New(&redis.Config{
		Host:     config.Redis.Host,
		Port:     config.Redis.Port,
		Password: config.Redis.Pw,
		DB:       config.Redis.Db,
	})
	app := service.NewApplication(service.Config{
		App:         config,
		EventEngine: eventEngine,
		Mongo:       notifyMongo,
		Validator:   valid,
	})
	tknSrv := token.New(token.Config{
		Expiration:     config.TokenSrv.Expiration,
		PublicKeyFile:  config.RSA.PublicKeyFile,
		PrivateKeyFile: config.RSA.PrivateKeyFile,
	})
	session := session.NewSessionApp(session.Config{
		Redis:       redis,
		EventEngine: eventEngine,
		Topic:       config.Session.Topic,
		TokenSrv:    tknSrv,
		Project:     config.TokenSrv.Project,
	})
	delivery := delivery.New(delivery.Config{
		App:         app,
		Config:      config,
		I18n:        i18n,
		Ctx:         ctx,
		EventEngine: eventEngine,
		TokenSrv:    tknSrv,
		SessionSrv:  session.Service,
	})
	delivery.Load()
}

func loadNotifyMongo(config config.App) *mongo.DB {
	uri := mongo.CalcMongoUri(mongo.UriParams{
		Host:  config.DB.MongoNotify.Host,
		Port:  config.DB.MongoNotify.Port,
		User:  config.DB.MongoNotify.Username,
		Pass:  config.DB.MongoNotify.Password,
		Db:    config.DB.MongoNotify.Database,
		Query: config.DB.MongoNotify.Query,
	})
	fmt.Printf("Mongo URI: %s", uri)
	d, err := mongo.New(uri, config.DB.MongoNotify.Database)
	if err != nil {
		fmt.Printf("Error loading mongo: %s", err.Error())
		panic(err)
	}
	fmt.Printf("Mongo loaded: %s", uri)
	return d
}
