package res

import (
	"api.turistikrota.com/notify/src/app/query"
	"api.turistikrota.com/notify/src/domain/notify"
)

type GetAllByChannelResponse struct {
	List          []DetailResponse `json:"list"`
	Total         int64            `json:"total"`
	FilteredTotal int64            `json:"filteredTotal"`
	IsNext        bool             `json:"isNext"`
	IsPrev        bool             `json:"isPrev"`
}

func (r *response) GetAllByChannel(result *query.GetAllByChannelResult) *GetAllByChannelResponse {
	res := &GetAllByChannelResponse{
		Total:         result.Total,
		FilteredTotal: result.FilteredTotal,
		IsNext:        result.IsNext,
		IsPrev:        result.IsPrev,
	}
	res.List = res.mapList(result.List, r)
	return res
}

func (r *GetAllByChannelResponse) mapList(list []notify.DetailResult, response *response) []DetailResponse {
	res := make([]DetailResponse, 0)
	for _, v := range list {
		res = append(res, *response.Detail(&v))
	}
	return res
}
