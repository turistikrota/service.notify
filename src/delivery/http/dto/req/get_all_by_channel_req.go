package req

import (
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.notify/src/app/query"
	"github.com/turistikrota/service.notify/src/domain/notify"
)

type GetAllByChannelRequest struct {
	PaginationRequest
	ListRequest
	Channel string `query:"channel" validate:"required"`
}

func (r *GetAllByChannelRequest) calcOffset() int64 {
	l := *r.Limit
	p := *r.Page
	offset := (p - 1) * l
	return offset
}

func (r *GetAllByChannelRequest) ToQuery() query.GetAllByChannelQuery {
	return query.GetAllByChannelQuery{
		Channel:   notify.Channel(r.Channel),
		Offset:    r.calcOffset(),
		Limit:     *r.Limit,
		StartDate: *r.StartDate,
		EndDate:   *r.EndDate,
	}
}

func (r *GetAllByChannelRequest) Default() *i18np.Error {
	r.PaginationRequest.Default()
	return r.ListRequest.Default()
}

func (r *request) GetAllByChannel() *GetAllByChannelRequest {
	return &GetAllByChannelRequest{
		PaginationRequest: *r.PaginationRequest(),
		ListRequest:       *r.ListRequest(),
	}
}
