package query

import (
	"context"

	"api.turistikrota.com/notify/src/domain/notify"
	"github.com/turistikrota/service.shared/decorator"
	"github.com/mixarchitecture/i18np"
)

type GetByUUIDQuery struct {
	UUID string
}

type GetByUUIDResult struct {
	Notify notify.DetailResult
}

type GetByUUIDHandler decorator.QueryHandler[GetByUUIDQuery, *GetByUUIDResult]

type getByUUIDHandler struct {
	notifyRepo notify.Repository
}

type GetByUUIDHandlerConfig struct {
	NotifyRepo notify.Repository
	CqrsBase   decorator.Base
}

func NewGetByUUIDHandler(config GetByUUIDHandlerConfig) GetByUUIDHandler {
	return decorator.ApplyQueryDecorators[GetByUUIDQuery, *GetByUUIDResult](
		getByUUIDHandler{
			notifyRepo: config.NotifyRepo,
		},
		config.CqrsBase,
	)
}

func (h getByUUIDHandler) Handle(ctx context.Context, query GetByUUIDQuery) (*GetByUUIDResult, *i18np.Error) {
	notify, err := h.notifyRepo.GetByUUID(ctx, query.UUID)
	if err != nil {
		return nil, err
	}
	return &GetByUUIDResult{Notify: *notify}, nil
}
