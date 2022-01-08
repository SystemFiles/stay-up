package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var App AppConfig

type AppConfig struct {
	AllowedOrigins []string
	DBHost string
	DBPort int64
	DBUser string
	DBPass string
	RefreshTimeMs int64
}

func (c *AppConfig) Init() error {
	godotenv.Load(".env")
	var dbHost, dbUser, dbPass string
	var dbPort, refreshTime int64
	var err error

	dbHost = os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbPort, err = strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
	if err != nil || (dbPort / 1000) < 1 {
		dbPort = 5432
	}

	dbUser = os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "stayup"
	}

	dbPass = os.Getenv("DB_PASS")
	if dbPass == "" {
		dbPass = "upstay"
	}

	refreshTime, err = strconv.ParseInt(os.Getenv("SERVICE_REFRESH_TIME_MS"), 10, 64)
	if err != nil || refreshTime < 1 {
		refreshTime = 10000 // 10s
	}

	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	if len(allowedOrigins) <= 1 || allowedOrigins == nil {
		allowedOrigins = []string{"*"}
	}

	App = AppConfig{
		AllowedOrigins: allowedOrigins,
		DBHost: dbHost,
		DBPort: dbPort,
		DBUser: dbUser,
		DBPass: dbPass,
		RefreshTimeMs: refreshTime,
	}

	return nil
}