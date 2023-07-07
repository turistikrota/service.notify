package query

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.notify/src/domain/notify"
	"github.com/turistikrota/service.shared/decorator"
)

type GetAllByRecipientQuery struct {
	Offset    int64
	Limit     int64
	StartDate time.Time
	EndDate   time.Time
	Recipient string
}

type GetAllByRecipientResult struct {
	List          []notify.DetailResult
	Total         int64
	FilteredTotal int64
	IsNext        bool
	IsPrev        bool
}

type GetAllByRecipientHandler decorator.QueryHandler[GetAllByRecipientQuery, *GetAllByRecipientResult]

type getAllByRecipientHandler struct {
	notifyRepo notify.Repository
}

type GetAllByRecipientHandlerConfig struct {
	NotifyRepo notify.Repository
	CqrsBase   decorator.Base
}

func NewGetAllByRecipientHandler(config GetAllByRecipientHandlerConfig) GetAllByRecipientHandler {
	return decorator.ApplyQueryDecorators[GetAllByRecipientQuery, *GetAllByRecipientResult](
		getAllByRecipientHandler{
			notifyRepo: config.NotifyRepo,
		},
		config.CqrsBase,
	)
}

func (h getAllByRecipientHandler) Handle(ctx context.Context, query GetAllByRecipientQuery) (*GetAllByRecipientResult, *i18np.Error) {
	res, err := h.notifyRepo.GetAllByRecipient(ctx, query.Recipient, notify.ListConfig{
		Offset:    query.Offset,
		Limit:     query.Limit,
		StartDate: query.StartDate,
		EndDate:   query.EndDate,
	})
	if err != nil {
		return nil, err
	}
	return &GetAllByRecipientResult{
		List:          res.List,
		Total:         res.Total,
		FilteredTotal: res.FilteredTotal,
		IsNext:        res.IsNext,
		IsPrev:        res.IsPrev,
	}, nil
}
