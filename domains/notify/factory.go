package notify

type Factory struct {
	Errors   Errors
	Messages Messages
}

func NewFactory() Factory {
	return Factory{
		Errors:   newErrors(),
		Messages: i18nMessages,
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}
