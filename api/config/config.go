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
	RDBHost string
	RDBPort string
	RDBPass string
	RefreshTimeMs int64
}

func (c *AppConfig) Init() error {
	godotenv.Load(".env")
	var rdbHost, rdbPort, rdbPass string
	var refreshTime int64
	var err error

	rdbHost = os.Getenv("REDIS_RDB_HOST")
	if rdbHost == "" {
		rdbHost = "localhost"
	}

	rdbPort = os.Getenv("REDIS_RDB_PORT")
	if rdbPort == "" {
		rdbPort = "5432"
	}

	rdbPass = os.Getenv("REDIS_RDB_PASS")

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
		RDBHost: rdbHost,
		RDBPort: rdbPort,
		RDBPass: rdbPass,
		RefreshTimeMs: refreshTime,
	}

	return nil
}