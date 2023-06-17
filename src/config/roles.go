package config

import "github.com/turistikrota/service.shared/base_roles"

type roles struct {
	base_roles.Roles
	ViewNotifyLogsByRecipient string
	ViewNotifyLogsByChannel   string
	ViewNotifyLogDetail       string
}

var Roles = roles{
	Roles:                     base_roles.BaseRoles,
	ViewNotifyLogsByRecipient: "view_notify_logs_by_recipient",
	ViewNotifyLogsByChannel:   "view_notify_logs_by_channel",
	ViewNotifyLogDetail:       "view_notify_log_detail",
}
