package req

type Request interface {
	GetByUUID() *GetByUUIDRequest
	GetAllByRecipient() *GetAllByRecipientRequest
	GetAllByChannel() *GetAllByChannelRequest
	PaginationRequest() *PaginationRequest
	ListRequest() *ListRequest
}

type request struct{}

func New() Request {
	return &request{}
}
