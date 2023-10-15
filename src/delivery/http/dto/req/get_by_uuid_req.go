package req

import "github.com/turistikrota/service.notify/src/app/query"

type GetByUUIDRequest struct {
	UUID string `param:"uuid" validate:"required,object_id"`
}

func (r *GetByUUIDRequest) ToQuery() query.GetByUUIDQuery {
	return query.GetByUUIDQuery{
		UUID: r.UUID,
	}
}

func (r *request) GetByUUID() *GetByUUIDRequest {
	return &GetByUUIDRequest{}
}
