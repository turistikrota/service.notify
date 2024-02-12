package config

import "github.com/turistikrota/service.shared/base_roles"

type actorConfigRoles struct {
	AddMail        string
	AddSms         string
	AddTelegram    string
	RemoveMail     string
	RemoveSms      string
	RemoveTelegram string
	EditMail       string
	EditSms        string
	EditTelegram   string
	View           string
	List           string
	Super          string
}

type notifyRoles struct {
	TestEmail string
	TestSms   string
}

type businessRoles struct {
	Super string
}

type roles struct {
	base_roles.Roles
	ActorConfig actorConfigRoles
	Business    businessRoles
	Notify      notifyRoles
}

var Roles = roles{
	Roles: base_roles.BaseRoles,
	ActorConfig: actorConfigRoles{
		AddMail:        "notify_actor_config.add_mail",
		AddSms:         "notify_actor_config.add_sms",
		AddTelegram:    "notify_actor_config.add_telegram",
		RemoveMail:     "notify_actor_config.remove_mail",
		RemoveSms:      "notify_actor_config.remove_sms",
		RemoveTelegram: "notify_actor_config.remove_telegram",
		EditMail:       "notify_actor_config.edit_mail",
		EditSms:        "notify_actor_config.edit_sms",
		EditTelegram:   "notify_actor_config.edit_telegram",
		View:           "notify_actor_config.view",
		List:           "notify_actor_config.list",
		Super:          "notify_actor_config.super",
	},
	Notify: notifyRoles{
		TestEmail: "notify.test_email",
		TestSms:   "notify.test_sms",
	},
	Business: businessRoles{
		Super: "business.super",
	},
}
