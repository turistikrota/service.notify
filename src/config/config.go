package config

type MongoNotify struct {
	Host       string `env:"MONGO_NOTIFY_HOST" envDefault:"localhost"`
	Port       string `env:"MONGO_NOTIFY_PORT" envDefault:"27017"`
	Username   string `env:"MONGO_NOTIFY_USERNAME" envDefault:""`
	Password   string `env:"MONGO_NOTIFY_PASSWORD" envDefault:""`
	Database   string `env:"MONGO_NOTIFY_DATABASE" envDefault:"empty"`
	Collection string `env:"MONGO_NOTIFY_COLLECTION" envDefault:"empties"`
	Query      string `env:"MONGO_NOTIFY_QUERY" envDefault:""`
}

type MailGoogle struct {
	FromMail string `env:"MAIL_GOOGLE_FROM_MAIL"`
	SmtpHost string `env:"MAIL_GOOGLE_SMTP_HOST"`
	SmtpPort string `env:"MAIL_GOOGLE_SMTP_PORT"`
	Password string `env:"MAIL_GOOGLE_PASSWORD"`
	Username string `env:"MAIL_GOOGLE_USERNAME"`
	Identity string `env:"MAIL_GOOGLE_IDENTITY"`
}

type RSA struct {
	PrivateKeyFile string `env:"RSA_PRIVATE_KEY"`
	PublicKeyFile  string `env:"RSA_PUBLIC_KEY"`
}

type Firebase struct {
	SecretFile string `env:"FIREBASE_SECRET_FILE"`
}

type Telegram struct {
	Token string `env:"TELEGRAM_TOKEN"`
}

type Redis struct {
	Host string `env:"REDIS_HOST"`
	Port string `env:"REDIS_PORT"`
	Pw   string `env:"REDIS_PASSWORD"`
	Db   int    `env:"REDIS_DB"`
}

type TokenSrv struct {
	Expiration int    `env:"TOKEN_EXPIRATION" envDefault:"3600"`
	Project    string `env:"TOKEN_PROJECT" envDefault:"empty"`
}

type NetGsm struct {
	UserName string `env:"NETGSM_USERNAME"`
	Password string `env:"NETGSM_PASSWORD"`
	Title    string `env:"NETGSM_TITLE"`
	Company  string `env:"NETGSM_COMPANY"`
}

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type Server struct {
	Host  string `env:"SERVER_HOST" envDefault:"localhost"`
	Port  int    `env:"SERVER_PORT" envDefault:"3000"`
	Group string `env:"SERVER_GROUP" envDefault:"auth"`
}

type HttpHeaders struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	Domain           string `env:"HTTP_HEADER_DOMAIN" envDefault:"*"`
}

type NotifyTopics struct {
	Email      string `env:"STREAMING_TOPIC_NOTIFY_EMAIL"`
	SMS        string `env:"STREAMING_TOPIC_NOTIFY_SMS"`
	Telegram   string `env:"STREAMING_TOPIC_NOTIFY_TELEGRAM"`
	Push       string `env:"STREAMING_TOPIC_NOTIFY_PUSH"`
	BatchEmail string `env:"STREAMING_TOPIC_NOTIFY_BATCH_EMAIL"`
	BatchSMS   string `env:"STREAMING_TOPIC_NOTIFY_BATCH_SMS"`
}

type Nats struct {
	Url     string   `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	Streams []string `env:"NATS_STREAMS" envDefault:""`
}

type Topics struct {
	Notify NotifyTopics
}

type Session struct {
	Topic string `env:"SESSION_TOPIC"`
}

type App struct {
	Protocol string `env:"PROTOCOL"`
	Adapters struct {
		Telegram   Telegram
		MailGoogle MailGoogle
		NetGsm     NetGsm
	}
	DB struct {
		MongoNotify MongoNotify
	}
	Server      Server
	HttpHeaders HttpHeaders
	I18n        I18n
	Topics      Topics
	Nats        Nats
	Redis       Redis
	TokenSrv    TokenSrv
	Session     Session
	RSA         RSA
	Firebase    Firebase
}
