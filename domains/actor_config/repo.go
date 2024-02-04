package actor_config

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WithActor struct {
	UUID string
	Name string
}

type Repository interface {
	GetByUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error)
	GetByUser(ctx context.Context, actor WithActor) (*Entity, *i18np.Error)
	GetByBusiness(ctx context.Context, actor WithActor) (*Entity, *i18np.Error)
	GetByUserUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error)
	GetByBusinessUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error)
	Filter(ctx context.Context, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error)

	AddTelegram(ctx context.Context, actor Actor, credential TelegramCredential) *i18np.Error
	AddMail(ctx context.Context, actor Actor, credential MailCredential) *i18np.Error
	AddSMS(ctx context.Context, actor Actor, credential SMSCredential) *i18np.Error

	UpdateTelegram(ctx context.Context, actor Actor, credential TelegramCredential) *i18np.Error
	UpdateMail(ctx context.Context, actor Actor, credential MailCredential) *i18np.Error
	UpdateSMS(ctx context.Context, actor Actor, credential SMSCredential) *i18np.Error

	RemoveTelegram(ctx context.Context, actor Actor, credentialName string) *i18np.Error
	RemoveMail(ctx context.Context, actor Actor, credentialName string) *i18np.Error
	RemoveSMS(ctx context.Context, actor Actor, credentialName string) *i18np.Error

	Create(ctx context.Context, entity *Entity) *i18np.Error
}

type repo struct {
	factory    Factory
	collection *mongo.Collection
	helper     mongo2.Helper[*Entity, *Entity]
}

func (r *repo) AddMail(ctx context.Context, actor Actor, credential MailCredential) *i18np.Error {
	filter := bson.M{
		actorField(actorFields.UUID): actor.UUID,
		actorField(actorFields.Name): actor.Name,
		actorField(actorFields.Type): actor.Type,
	}
	update := bson.M{
		"$addToSet": bson.M{
			fields.Mail: credential,
		},
		"$set": bson.M{
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) AddSMS(ctx context.Context, actor Actor, credential SMSCredential) *i18np.Error {
	filter := bson.M{
		actorField(actorFields.UUID): actor.UUID,
		actorField(actorFields.Name): actor.Name,
		actorField(actorFields.Type): actor.Type,
	}
	update := bson.M{
		"$addToSet": bson.M{
			fields.SMS: credential,
		},
		"$set": bson.M{
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) AddTelegram(ctx context.Context, actor Actor, credential TelegramCredential) *i18np.Error {
	filter := bson.M{
		actorField(actorFields.UUID): actor.UUID,
		actorField(actorFields.Name): actor.Name,
		actorField(actorFields.Type): actor.Type,
	}
	update := bson.M{
		"$addToSet": bson.M{
			fields.Telegram: credential,
		},
		"$set": bson.M{
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Create(ctx context.Context, e *Entity) *i18np.Error {
	_, err := r.collection.InsertOne(ctx, e)
	if err != nil {
		return r.factory.Errors.Failed("create")
	}
	return nil
}

// Filter implements Repository.
func (*repo) Filter(ctx context.Context, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error) {
	panic("unimplemented")
}

// GetByBusiness implements Repository.
func (*repo) GetByBusiness(ctx context.Context, actor WithActor) (*Entity, *i18np.Error) {
	panic("unimplemented")
}

// GetByBusinessUUID implements Repository.
func (*repo) GetByBusinessUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error) {
	panic("unimplemented")
}

// GetByUUID implements Repository.
func (*repo) GetByUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error) {
	panic("unimplemented")
}

// GetByUser implements Repository.
func (*repo) GetByUser(ctx context.Context, actor WithActor) (*Entity, *i18np.Error) {
	panic("unimplemented")
}

// GetByUserUUID implements Repository.
func (*repo) GetByUserUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error) {
	panic("unimplemented")
}

// RemoveMail implements Repository.
func (*repo) RemoveMail(ctx context.Context, actor Actor, credentialName string) *i18np.Error {
	panic("unimplemented")
}

// RemoveSMS implements Repository.
func (*repo) RemoveSMS(ctx context.Context, actor Actor, credentialName string) *i18np.Error {
	panic("unimplemented")
}

// RemoveTelegram implements Repository.
func (*repo) RemoveTelegram(ctx context.Context, actor Actor, credentialName string) *i18np.Error {
	panic("unimplemented")
}

// UpdateMail implements Repository.
func (*repo) UpdateMail(ctx context.Context, actor Actor, credential MailCredential) *i18np.Error {
	panic("unimplemented")
}

// UpdateSMS implements Repository.
func (*repo) UpdateSMS(ctx context.Context, actor Actor, credential SMSCredential) *i18np.Error {
	panic("unimplemented")
}

// UpdateTelegram implements Repository.
func (*repo) UpdateTelegram(ctx context.Context, actor Actor, credential TelegramCredential) *i18np.Error {
	panic("unimplemented")
}

func NewRepo(collection *mongo.Collection, factory Factory) Repository {
	return &repo{
		factory:    factory,
		collection: collection,
		helper:     mongo2.NewHelper[*Entity, *Entity](collection, createEntity),
	}
}

func createEntity() **Entity {
	return new(*Entity)
}
