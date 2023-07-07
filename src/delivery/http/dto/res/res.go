package res

import (
	"github.com/turistikrota/service.notify/src/app/query"
	"github.com/turistikrota/service.notify/src/domain/notify"
)

type Response interface {
	Detail(notify *notify.DetailResult) *DetailResponse
	GetAllByRecipient(res *query.GetAllByRecipientResult) *GetAllByRecipientResponse
	GetAllByChannel(notify *query.GetAllByChannelResult) *GetAllByChannelResponse
}

type response struct{}

func New() Response {
	return &response{}
}
