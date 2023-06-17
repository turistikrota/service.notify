package res

import (
	"api.turistikrota.com/notify/src/app/query"
	"api.turistikrota.com/notify/src/domain/notify"
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
