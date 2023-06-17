package http

type successMessages struct {
	GetByUUID         string
	GetAllByRecipient string
	GetAllByChannel   string
}

type errorMessages struct{}

type messages struct {
	Success successMessages
	Error   errorMessages
}

var Messages = messages{
	Success: successMessages{
		GetByUUID:         "http_success_get_by_uuid",
		GetAllByRecipient: "http_success_get_all_by_recipient",
		GetAllByChannel:   "http_success_get_all_by_channel",
	},
	Error: errorMessages{},
}
