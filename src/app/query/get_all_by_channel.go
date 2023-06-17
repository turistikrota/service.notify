package query

import (
	"context"
	"time"

	"api.turistikrota.com/notify/src/domain/notify"
	"github.com/turistikrota/service.shared/decorator"
	"github.com/mixarchitecture/i18np"
)

type GetAllByChannelQuery struct {
	Offset    int64
	Limit     int64
	StartDate time.Time
	EndDate   time.Time
	Channel   notify.Channel
}

type GetAllByChannelResult struct {
	List          []notify.DetailResult
	Total         int64
	FilteredTotal int64
	IsNext        bool
	IsPrev        bool
}

type GetAllByChannelHandler decorator.QueryHandler[GetAllByChannelQuery, *GetAllByChannelResult]

type getAllByChannelHandler struct {
	notifyRepo notify.Repository
}

type GetAllByChannelHandlerConfig struct {
	NotifyRepo notify.Repository
	CqrsBase   decorator.Base
}

func NewGetAllByChannelHandler(config GetAllByChannelHandlerConfig) GetAllByChannelHandler {
	return decorator.ApplyQueryDecorators[GetAllByChannelQuery, *GetAllByChannelResult](
		getAllByChannelHandler{
			notifyRepo: config.NotifyRepo,
		},
		config.CqrsBase,
	)
}

func (h getAllByChannelHandler) Handle(ctx context.Context, query GetAllByChannelQuery) (*GetAllByChannelResult, *i18np.Error) {
	res, err := h.notifyRepo.GetAllByChannel(ctx, query.Channel, notify.ListConfig{
		Offset:    query.Offset,
		Limit:     query.Limit,
		StartDate: query.StartDate,
		EndDate:   query.EndDate,
	})
	if err != nil {
		return nil, err
	}
	return &GetAllByChannelResult{
		List:          res.List,
		Total:         res.Total,
		FilteredTotal: res.FilteredTotal,
		IsNext:        res.IsNext,
		IsPrev:        res.IsPrev,
	}, nil
}
