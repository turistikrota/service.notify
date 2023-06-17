package req

import (
	"api.turistikrota.com/notify/src/app/query"
	"github.com/mixarchitecture/i18np"
)

type GetAllByRecipientRequest struct {
	PaginationRequest
	ListRequest
	Recipient string `query:"recipient" validate:"required"`
}

func (r *GetAllByRecipientRequest) calcOffset() int64 {
	l := *r.Limit
	p := *r.Page
	offset := (p - 1) * l
	return offset
}

func (r *GetAllByRecipientRequest) ToQuery() query.GetAllByRecipientQuery {
	return query.GetAllByRecipientQuery{
		Recipient: r.Recipient,
		Offset:    r.calcOffset(),
		Limit:     *r.Limit,
		StartDate: *r.StartDate,
		EndDate:   *r.EndDate,
	}
}

func (r *GetAllByRecipientRequest) Default() *i18np.Error {
	r.PaginationRequest.Default()
	return r.ListRequest.Default()
}

func (r *request) GetAllByRecipient() *GetAllByRecipientRequest {
	return &GetAllByRecipientRequest{
		PaginationRequest: *r.PaginationRequest(),
		ListRequest:       *r.ListRequest(),
	}
}
