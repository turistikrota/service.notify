package actor_config

type messages struct {
	Failed      string
	NotFound    string
	InvalidUUID string
}

var i18nMessages = messages{
	Failed:      "actor_config_failed",
	NotFound:    "actor_config_not_found",
	InvalidUUID: "actor_config_invalid_uuid",
}
