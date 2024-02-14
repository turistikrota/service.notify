package actor_config

type fieldsType struct {
	UUID      string
	Actor     string
	Telegram  string
	Mail      string
	SMS       string
	UpdatedAt string
}

type actorFieldsType struct {
	UUID string
	Name string
	Type string
}

type telegramFieldsType struct {
	Name   string
	ChatID string
}

type mailFieldsType struct {
	Name  string
	Email string
}

type smsFieldsType struct {
	Name        string
	Phone       string
	CountryCode string
}

var fields = fieldsType{
	UUID:      "_id",
	Actor:     "actor",
	Telegram:  "telegram",
	Mail:      "mail",
	SMS:       "sms",
	UpdatedAt: "updated_at",
}

var telegramFields = telegramFieldsType{
	Name:   "name",
	ChatID: "chat_id",
}

var mailFields = mailFieldsType{
	Name:  "name",
	Email: "email",
}

var smsFields = smsFieldsType{
	Name:        "name",
	Phone:       "phone",
	CountryCode: "country_code",
}

var actorFields = actorFieldsType{
	UUID: "uuid",
	Name: "name",
	Type: "type",
}

func actorField(field string) string {
	return fields.Actor + "." + field
}

func telegramField(field string) string {
	return fields.Telegram + "." + field
}

func mailField(field string) string {
	return fields.Mail + "." + field
}

func smsField(field string) string {
	return fields.SMS + "." + field
}

func telegramFieldInArray(field string) string {
	return fields.Telegram + ".$." + field
}

func mailFieldInArray(field string) string {
	return fields.Mail + ".$." + field
}

func smsFieldInArray(field string) string {
	return fields.SMS + ".$." + field
}
