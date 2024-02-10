package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.notify/domains/actor_config"
	"github.com/turistikrota/service.notify/pkg/utils"
)

type ActorConfigFilterQuery struct {
	*actor_config.FilterEntity
	*utils.Pagination
}

type ActorConfigFilterRes struct {
	List *list.Result[actor_config.AdminListDto]
}

type ActorConfigFilterHandler cqrs.HandlerFunc[ActorConfigFilterQuery, *ActorConfigFilterRes]

func NewActorConfigFilterHandler(factory actor_config.Factory, repo actor_config.Repository) ActorConfigFilterHandler {
	return func(ctx context.Context, query ActorConfigFilterQuery) (*ActorConfigFilterRes, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		res, err := repo.Filter(ctx, *query.FilterEntity, list.Config{
			Offset: offset,
			Limit:  *query.Limit,
		})
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		result := &list.Result[actor_config.AdminListDto]{
			Total:         res.Total,
			List:          make([]actor_config.AdminListDto, 0, len(res.List)),
			FilteredTotal: res.FilteredTotal,
			Page:          res.Page,
			IsNext:        res.IsNext,
			IsPrev:        res.IsPrev,
		}
		for _, item := range res.List {
			result.List = append(result.List, *item.ToAdminList())
		}
		return &ActorConfigFilterRes{
			List: result,
		}, nil
	}
}
