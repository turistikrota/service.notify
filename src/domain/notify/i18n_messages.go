package notify

type messages struct {
	Failed           string
	TypeNotFound     string
	ValidationFailed string
	NotFound         string
}

var I18nMessages = messages{
	Failed:           "error_notify_failed",
	TypeNotFound:     "error_notify_type_not_found",
	ValidationFailed: "error_notify_validation_failed",
	NotFound:         "error_notify_not_found",
}
