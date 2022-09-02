package config

type DBConfig struct {
	DBPort         int
	DBUrl          string
	DBjaeger_url   string
	DBsentry_url   string
	DBkafka_broker string
	DBsome_app_id  string
	DBsome_app_key string
}

func Config() DBConfig {
	return DBConfig{}
}