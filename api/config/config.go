package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var App AppConfig

type AppConfig struct {
	DBHost string
	DBPort int64
	DBUser string
	DBPass string
	DefaultServiceRefreshTimeMs int64
}

func (c *AppConfig) Init() error {
	godotenv.Load()
	var dbHost, dbUser, dbPass string
	var dbPort, refreshTime int64
	var err error

	if os.Getenv("DB_HOST") == "" {
		dbHost = "localhost"
	}

	dbPort, err = strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
	if err != nil || (dbPort / 1000) < 1 {
		// set default value
		dbPort = 5432
	}

	if os.Getenv("DB_USER") == "" {
		dbUser = "stayup"
	}

	if os.Getenv("DB_PASS") == "" {
		dbPass = "upstay"
	}

	refreshTime, err = strconv.ParseInt(os.Getenv("SERVICE_REFRESH_TIME_MS"), 10, 64)
	if err != nil || refreshTime < 1 {
		// set default value
		refreshTime = 10000 // 10s
	}

	App = AppConfig{
		DBHost: dbHost,
		DBPort: dbPort,
		DBUser: dbUser,
		DBPass: dbPass,
		DefaultServiceRefreshTimeMs: refreshTime,
	}

	return nil
}