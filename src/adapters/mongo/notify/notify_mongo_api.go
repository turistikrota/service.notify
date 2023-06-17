package notify

import (
	"context"

	"api.turistikrota.com/notify/src/adapters/mongo/notify/entity"
	"api.turistikrota.com/notify/src/domain/notify"
	shared_mongo "github.com/turistikrota/service.shared/db/mongo"
	"github.com/mixarchitecture/i18np"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) Log(ctx context.Context, notify *notify.Notify, data interface{}) *i18np.Error {
	n := &entity.MongoNotify{}
	res, err := r.collection.InsertOne(ctx, n.FromNotify(notify, data))
	if err != nil {
		return r.notifyFactory.Errors.Failed("log")
	}
	notify.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r *repo) GetByUUID(ctx context.Context, uuid string) (*notify.DetailResult, *i18np.Error) {
	n := &entity.MongoNotify{}
	oid, error := shared_mongo.TransformId(uuid)
	if error != nil {
		return nil, r.notifyFactory.Errors.Failed("parse id")
	}
	res := r.collection.FindOne(ctx, bson.M{"_id": oid})
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, r.notifyFactory.Errors.NotFound()
		}
		return nil, r.notifyFactory.Errors.Failed("get")
	}
	err := res.Decode(n)
	if err != nil {
		return nil, r.notifyFactory.Errors.Failed("get")
	}
	detail := n.ToDetail()
	return detail, nil
}

func (r *repo) GetAllByRecipient(ctx context.Context, recipient string, config notify.ListConfig) (*notify.ListResult, *i18np.Error) {
	return r.getAllWithFilter(ctx, bson.M{
		"recipient": recipient,
		"created_at": bson.M{
			"$gte": config.StartDate,
			"$lte": config.EndDate,
		},
	}, config)
}

func (r *repo) GetAllByChannel(ctx context.Context, ch notify.Channel, config notify.ListConfig) (*notify.ListResult, *i18np.Error) {
	return r.getAllWithFilter(ctx, bson.M{
		"type": ch,
		"created_at": bson.M{
			"$gte": config.StartDate,
			"$lte": config.EndDate,
		},
	}, config)
}

func (r *repo) getAllWithFilter(ctx context.Context, filter interface{}, config notify.ListConfig) (*notify.ListResult, *i18np.Error) {
	n := []notify.DetailResult{}
	cur, error := r.collection.Find(ctx, filter, options.Find().SetLimit(config.Limit).SetSkip(config.Offset))
	if error != nil {
		return nil, r.notifyFactory.Errors.Failed("get 1")
	}
	for cur.Next(ctx) {
		var m *entity.MongoNotify
		err := cur.Decode(&m)
		if err != nil {
			return nil, r.notifyFactory.Errors.Failed("get 2")
		}
		n = append(n, *m.ToDetail())
	}
	filteredTotal, err := r.getCountWithFilter(ctx, filter, config)
	if err != nil {
		return nil, r.notifyFactory.Errors.Failed("get 3")
	}
	total, err := r.getCountWithFilter(ctx, bson.M{}, config)
	if err != nil {
		return nil, err
	}
	return r.toResult(n, total, filteredTotal, config), nil
}

func (r *repo) getCountWithFilter(ctx context.Context, filter interface{}, config notify.ListConfig) (int64, *i18np.Error) {
	total, err := r.collection.CountDocuments(ctx, filter, options.Count().SetLimit(config.Limit).SetSkip(config.Offset))
	if err != nil {
		return 0, r.notifyFactory.Errors.Failed("get")
	}
	return total, nil
}

func (r *repo) toResult(list []notify.DetailResult, total int64, filteredTotal int64, config notify.ListConfig) *notify.ListResult {
	return &notify.ListResult{
		List:          list,
		Total:         total,
		FilteredTotal: filteredTotal,
		IsNext:        ((config.Offset) + config.Limit) < total,
		IsPrev:        config.Offset > 0,
	}
}
