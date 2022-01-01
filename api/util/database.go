package util

import (
	"fmt"
	"log"
	"sync"

	"github.com/systemfiles/stay-up/api/config"
	"github.com/systemfiles/stay-up/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseInstance *gorm.DB

var (
	instance DatabaseInstance
	once sync.Once
)

func GetDBInstance() (DatabaseInstance, error) {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=stayup_db port=%d sslmode=disable TimeZone=America/Chicago", config.App.DBHost, config.App.DBUser, config.App.DBPass, config.App.DBPort)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to start database connection. reason: %s", err.Error())
		}

		// configure db instance
		db.AutoMigrate(&models.Service{})

		instance = DatabaseInstance(db)
	})

	return instance, nil
}