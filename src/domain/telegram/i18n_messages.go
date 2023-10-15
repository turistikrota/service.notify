package telegram

type messages struct {
	Failed           string
	ValidationFailed string
}

var I18nMessages = messages{
	Failed:           "error_notify_failed",
	ValidationFailed: "error_notify_validation_failed",
}
