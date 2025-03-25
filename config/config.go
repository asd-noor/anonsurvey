package config

import (
	"anonsurvey/logger"
	"encoding/json"
	"io"
	"log/slog"
	"os"
)

var conf config

func Config() config {
	return conf
}

func init() {
	logWriters := []io.Writer{os.Stderr}

	configPath := os.Getenv("APP_CONFIG")
	if configPath == "" {
		hardCodedLocal()
		return
	}

	cb, err := os.ReadFile(configPath)
	if err != nil {
		logger.Init(slog.LevelDebug)
		logger.Fatal("cannot read config file", "error", err, "filepath", configPath)
	}

	c := config{}
	if err := json.Unmarshal(cb, &c); err != nil {
		logger.Init(slog.LevelDebug)
		logger.Fatal("cannot unmarshal config from json", "error", err, "filepath", configPath)
	}

	conf = c
	logger.Init(slog.Level(c.App.LogLevel), logWriters...)
}

func hardCodedLocal() {
	conf = config{
		App: AppConfig{
			AppEnv:   "LOCAL",
			Port:     8080,
			LogLevel: -4,
		},
		DB: DBConfig{
			Host:   "localhost",
			Port:   3306,
			User:   "anonsurvey",
			Pass:   "yevrusnona",
			Schema: "dbsurvey",
			Debug:  true,
		},
	}

	logger.Init(slog.Level(conf.App.LogLevel))
	logger.Warn("LOADED HARD CODED CONFIG", "config", conf)
}
