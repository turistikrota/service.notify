package res

import (
	"github.com/turistikrota/service.notify/src/app/query"
	"github.com/turistikrota/service.notify/src/domain/notify"
)

type GetAllByRecipientResponse struct {
	List          []DetailResponse `json:"list"`
	Total         int64            `json:"total"`
	FilteredTotal int64            `json:"filteredTotal"`
	IsNext        bool             `json:"isNext"`
	IsPrev        bool             `json:"isPrev"`
}

func (r *response) GetAllByRecipient(result *query.GetAllByRecipientResult) *GetAllByRecipientResponse {
	res := &GetAllByRecipientResponse{
		Total:         result.Total,
		FilteredTotal: result.FilteredTotal,
		IsNext:        result.IsNext,
		IsPrev:        result.IsPrev,
	}
	res.List = res.mapList(result.List, r)
	return res
}

func (r *GetAllByRecipientResponse) mapList(list []notify.DetailResult, response *response) []DetailResponse {
	res := make([]DetailResponse, 0)
	for _, v := range list {
		res = append(res, *response.Detail(&v))
	}
	return res
}
