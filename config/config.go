package config

type MongoActorConfig struct {
	Host       string `env:"MONGO_ACTOR_CONFIG_HOST" envDefault:"localhost"`
	Port       string `env:"MONGO_ACTOR_CONFIG_PORT" envDefault:"27017"`
	Username   string `env:"MONGO_ACTOR_CONFIG_USERNAME" envDefault:""`
	Password   string `env:"MONGO_ACTOR_CONFIG_PASSWORD" envDefault:""`
	Database   string `env:"MONGO_ACTOR_CONFIG_DATABASE" envDefault:"account"`
	Collection string `env:"MONGO_ACTOR_CONFIG_COLLECTION" envDefault:"accounts"`
	Query      string `env:"MONGO_ACTOR_CONFIG_QUERY" envDefault:""`
}

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type Http struct {
	Host  string `env:"SERVER_HOST" envDefault:"localhost"`
	Port  int    `env:"SERVER_PORT" envDefault:"3000"`
	Group string `env:"SERVER_GROUP" envDefault:"account"`
}

type Firebase struct {
	SecretFile string `env:"FIREBASE_SECRET_FILE"`
}

type NetGsm struct {
	UserName string `env:"NETGSM_USERNAME"`
	Password string `env:"NETGSM_PASSWORD"`
	Title    string `env:"NETGSM_TITLE"`
	Company  string `env:"NETGSM_COMPANY"`
}

type MailGun struct {
	FromMail string `env:"MAIL_GUN_FROM_MAIL"`
	SmtpHost string `env:"MAIL_GUN_SMTP_HOST"`
	SmtpPort string `env:"MAIL_GUN_SMTP_PORT"`
	Password string `env:"MAIL_GUN_PASSWORD"`
	Username string `env:"MAIL_GUN_USERNAME"`
	Identity string `env:"MAIL_GUN_IDENTITY"`
}

type Telegram struct {
	Token string `env:"TELEGRAM_TOKEN"`
}

type Grpc struct {
	Port int `env:"GRPC_PORT" envDefault:"3001"`
}

type Redis struct {
	Host string `env:"REDIS_HOST"`
	Port string `env:"REDIS_PORT"`
	Pw   string `env:"REDIS_PASSWORD"`
	Db   int    `env:"REDIS_DB"`
}

type CacheRedis struct {
	Host string `env:"REDIS_CACHE_HOST"`
	Port string `env:"REDIS_CACHE_PORT"`
	Pw   string `env:"REDIS_CACHE_PASSWORD"`
	Db   int    `env:"REDIS_CACHE_DB"`
}

type HttpHeaders struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	Domain           string `env:"HTTP_HEADER_DOMAIN" envDefault:"*"`
}

type TokenSrv struct {
	Expiration int    `env:"TOKEN_EXPIRATION" envDefault:"3600"`
	Project    string `env:"TOKEN_PROJECT" envDefault:"empty"`
}

type Session struct {
	Topic string `env:"SESSION_TOPIC"`
}

type RSA struct {
	PrivateKeyFile string `env:"RSA_PRIVATE_KEY"`
	PublicKeyFile  string `env:"RSA_PUBLIC_KEY"`
}

type Topics struct {
}

type Nats struct {
	Url     string   `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	Streams []string `env:"NATS_STREAMS" envDefault:""`
}

type App struct {
	Protocol string `env:"PROTOCOL" envDefault:"http"`
	DB       struct {
		ActorConfig MongoActorConfig
	}
	Adapters struct {
		MailGun  MailGun
		Telegram Telegram
		NetGsm   NetGsm
		Firebase Firebase
	}
	Grpc        Grpc
	Http        Http
	HttpHeaders HttpHeaders
	I18n        I18n
	Topics      Topics
	Session     Session
	Nats        Nats
	Redis       Redis
	TokenSrv    TokenSrv
	CacheRedis  CacheRedis
	RSA         RSA
	Firebase    Firebase
}
