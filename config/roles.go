package config

import "github.com/turistikrota/service.shared/base_roles"

type actorConfigRoles struct{}

type businessRoles struct {
	Super string
}

type roles struct {
	base_roles.Roles
	ActorConfig actorConfigRoles
	Business    businessRoles
}

var Roles = roles{
	Roles:       base_roles.BaseRoles,
	ActorConfig: actorConfigRoles{},
	Business: businessRoles{
		Super: "business.super",
	},
}
