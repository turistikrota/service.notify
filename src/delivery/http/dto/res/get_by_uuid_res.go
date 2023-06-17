package res

import (
	"api.turistikrota.com/notify/src/domain/notify"
	"github.com/turistikrota/service.shared/formats"
)

type DetailResponse struct {
	UUID      string      `json:"uuid"`
	Channel   string      `json:"channel"`
	Recipient string      `json:"recipient"`
	CreatedAt string      `json:"createdAt"`
	UpdatedAt string      `json:"updatedAt"`
	Details   interface{} `json:"details"`
}

func (r *response) Detail(notify *notify.DetailResult) *DetailResponse {
	return &DetailResponse{
		UUID:      notify.UUID,
		Channel:   string(notify.Type),
		Recipient: notify.Recipient,
		CreatedAt: notify.CreatedAt.Format(formats.ISO8601),
		UpdatedAt: notify.UpdatedAt.Format(formats.ISO8601),
		Details:   notify.Data,
	}
}
