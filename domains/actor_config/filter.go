package actor_config

import "go.mongodb.org/mongo-driver/bson"

type FilterEntity struct{}

func (r *repo) filterToBson(filter FilterEntity) bson.M {
	list := make([]bson.M, 0)
	listLen := len(list)
	if listLen == 0 {
		return bson.M{}
	}
	if listLen == 1 {
		return list[0]
	}
	return bson.M{
		"$and": list,
	}
}
