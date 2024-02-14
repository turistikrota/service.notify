package notify

type Messages struct {
	Failed            string
	NotMailConfigured string
	NotSmsConfigured  string

	TestMailSubject string
	TestMailContent string

	TestSmsContent string

	TestTelegramContent string
}

var i18nMessages = Messages{
	Failed:              "notify_failed",
	NotMailConfigured:   "notify_not_mail_configured",
	NotSmsConfigured:    "notify_not_sms_configured",
	TestMailSubject:     "notify_test_mail_subject",
	TestMailContent:     "notify_test_mail_content",
	TestSmsContent:      "notify_test_sms_content",
	TestTelegramContent: "notify_test_telegram_content",
}
