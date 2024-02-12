package actor_config

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WithActor struct {
	UUID string
	Name string
}

type Repository interface {
	GetByUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error)
	GetByUser(ctx context.Context, actor WithActor) (*Entity, *i18np.Error)
	GetByBusiness(ctx context.Context, actor WithActor) (*Entity, *i18np.Error)
	GetByUserName(ctx context.Context, name string) (*Entity, *i18np.Error)
	GetByBusinessUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error)
	Filter(ctx context.Context, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error)

	AddTelegram(ctx context.Context, actor Actor, credential TelegramCredential) *i18np.Error
	AddMail(ctx context.Context, actor Actor, credential MailCredential) *i18np.Error
	AddSMS(ctx context.Context, actor Actor, credential SMSCredential) *i18np.Error

	UpdateTelegram(ctx context.Context, actor Actor, credential TelegramCredential, oldName string) *i18np.Error
	UpdateMail(ctx context.Context, actor Actor, credential MailCredential, oldName string) *i18np.Error
	UpdateSMS(ctx context.Context, actor Actor, credential SMSCredential, oldName string) *i18np.Error

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

func (r *repo) Filter(ctx context.Context, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error) {
	filters := r.filterToBson(filter)
	l, err := r.helper.GetListFilter(ctx, filters, r.sort(r.filterOptions(listConfig), filter))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, filters)
	if _err != nil {
		return nil, _err
	}
	total, _err := r.helper.GetFilterCount(ctx, bson.M{})
	if _err != nil {
		return nil, _err
	}
	return &list.Result[*Entity]{
		IsNext:        filtered > listConfig.Offset+listConfig.Limit,
		IsPrev:        listConfig.Offset > 0,
		FilteredTotal: filtered,
		Total:         total,
		Page:          listConfig.Offset/listConfig.Limit + 1,
		List:          l,
	}, nil
}

func (r *repo) GetByBusiness(ctx context.Context, actor WithActor) (*Entity, *i18np.Error) {
	filter := bson.M{
		actorField(actorFields.UUID): actor.UUID,
		actorField(actorFields.Name): actor.Name,
		actorField(actorFields.Type): ActorTypeBusiness,
	}
	e, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, r.factory.Errors.Failed("get")
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return *e, nil
}

func (r *repo) GetByBusinessUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error) {
	filter := bson.M{
		actorField(actorFields.UUID): uuid,
		actorField(actorFields.Type): ActorTypeBusiness,
	}
	e, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, r.factory.Errors.Failed("get")
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return *e, nil
}

func (r *repo) GetByUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error) {
	id, _err := mongo2.TransformId(uuid)
	if _err != nil {
		return nil, r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	e, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, r.factory.Errors.Failed("get")
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return *e, nil
}

func (r *repo) GetByUser(ctx context.Context, actor WithActor) (*Entity, *i18np.Error) {
	filter := bson.M{
		actorField(actorFields.UUID): actor.UUID,
		actorField(actorFields.Name): actor.Name,
		actorField(actorFields.Type): ActorTypeUser,
	}
	e, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, r.factory.Errors.Failed("get")
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return *e, nil
}

func (r *repo) GetByUserName(ctx context.Context, name string) (*Entity, *i18np.Error) {
	filter := bson.M{
		actorField(actorFields.Name): name,
		actorField(actorFields.Type): ActorTypeUser,
	}
	e, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, r.factory.Errors.Failed("get")
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return *e, nil
}

func (r *repo) RemoveMail(ctx context.Context, actor Actor, credentialName string) *i18np.Error {
	filter := bson.M{
		actorField(actorFields.UUID): actor.UUID,
		actorField(actorFields.Name): actor.Name,
		actorField(actorFields.Type): actor.Type,
		mailField(mailFields.Name):   credentialName,
	}
	update := bson.M{
		"$pull": bson.M{
			fields.Mail: bson.M{
				mailFields.Name: credentialName,
			},
		},
		"$set": bson.M{
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) RemoveSMS(ctx context.Context, actor Actor, credentialName string) *i18np.Error {
	filter := bson.M{
		actorField(actorFields.UUID): actor.UUID,
		actorField(actorFields.Name): actor.Name,
		actorField(actorFields.Type): actor.Type,
		smsField(smsFields.Name):     credentialName,
	}
	update := bson.M{
		"$pull": bson.M{
			fields.SMS: bson.M{
				smsFields.Name: credentialName,
			},
		},
		"$set": bson.M{
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) RemoveTelegram(ctx context.Context, actor Actor, credentialName string) *i18np.Error {
	filter := bson.M{
		actorField(actorFields.UUID):       actor.UUID,
		actorField(actorFields.Name):       actor.Name,
		actorField(actorFields.Type):       actor.Type,
		telegramField(telegramFields.Name): credentialName,
	}
	update := bson.M{
		"$pull": bson.M{
			fields.Telegram: bson.M{
				telegramFields.Name: credentialName,
			},
		},
		"$set": bson.M{
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) UpdateMail(ctx context.Context, actor Actor, credential MailCredential, oldName string) *i18np.Error {
	filter := bson.M{
		actorField(actorFields.UUID): actor.UUID,
		actorField(actorFields.Name): actor.Name,
		actorField(actorFields.Type): actor.Type,
		mailField(mailFields.Name):   oldName,
	}
	update := bson.M{
		"$set": bson.M{
			mailFieldInArray(mailFields.Email): credential.Email,
			fields.UpdatedAt:                   time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) UpdateSMS(ctx context.Context, actor Actor, credential SMSCredential, oldName string) *i18np.Error {
	filter := bson.M{
		actorField(actorFields.UUID): actor.UUID,
		actorField(actorFields.Name): actor.Name,
		actorField(actorFields.Type): actor.Type,
		smsField(smsFields.Name):     oldName,
	}
	update := bson.M{
		"$set": bson.M{
			smsFieldInArray(smsFields.Phone):       credential.Phone,
			smsFieldInArray(smsFields.CountryCode): credential.CountryCode,
			fields.UpdatedAt:                       time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) UpdateTelegram(ctx context.Context, actor Actor, credential TelegramCredential, oldName string) *i18np.Error {
	filter := bson.M{
		actorField(actorFields.UUID):       actor.UUID,
		actorField(actorFields.Name):       actor.Name,
		actorField(actorFields.Type):       actor.Type,
		telegramField(telegramFields.Name): oldName,
	}
	update := bson.M{
		"$set": bson.M{
			telegramFieldInArray(telegramFields.ChatID): credential.ChatID,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
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

func (r *repo) filterOptions(listConfig list.Config) *options.FindOptions {
	opts := &options.FindOptions{}
	opts.SetSkip(listConfig.Offset).SetLimit(listConfig.Limit)
	return opts
}

func (r *repo) sort(opts *options.FindOptions, filter FilterEntity) *options.FindOptions {
	order := -1
	field := fields.UpdatedAt
	/*
		if filter.Order == OrderAsc {
			order = 1
		}
		switch filter.Sort {
		case SortByMostRecent:
			field = fields.UpdatedAt
		case SortByNearest:
			field = locationField(locationFields.Coordinates)
		case SortByPrice:
			field = priceField(priceFields.Price)
		}
	*/
	opts.SetSort(bson.D{{Key: field, Value: order}})
	return opts
}
