package config

import "github.com/turistikrota/service.shared/base_roles"

type actorConfigRoles struct {
	Add    string
	Remove string
	Edit   string
	View   string
	List   string
	Super  string
}

type notifyRoles struct {
	TestEmail    string
	TestSms      string
	TestTelegram string
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
		Add:    "notify_actor_config.add",
		Remove: "notify_actor_config.remove",
		Edit:   "notify_actor_config.edit",
		View:   "notify_actor_config.view",
		List:   "notify_actor_config.list",
		Super:  "notify_actor_config.super",
	},
	Notify: notifyRoles{
		TestEmail:    "notify.test_email",
		TestSms:      "notify.test_sms",
		TestTelegram: "notify.test_telegram",
	},
	Business: businessRoles{
		Super: "business.super",
	},
}
