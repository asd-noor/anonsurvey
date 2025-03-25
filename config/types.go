package config

type AppConfig struct {
	AppEnv   string
	LogLevel int
	Port     int
}

type DBConfig struct {
	Host   string
	Port   int
	User   string
	Pass   string
	Schema string
	Debug  bool
}

type config struct {
	App AppConfig
	DB  DBConfig
}
