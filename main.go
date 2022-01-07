package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/systemfiles/stay-up/api/config"
	"github.com/systemfiles/stay-up/api/controller"
	"github.com/systemfiles/stay-up/api/tasks"
)

// Custom struct validation
type CustomValidator struct {
	validator *validator.Validate
}

func (csv *CustomValidator) Validate(i interface{}) error {
	if err := csv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	// Create echo application
	e := echo.New()

	// Load app config
	err := config.App.Init()
	if err != nil {
		log.Fatal("Failed to get app configuration ... cannot start")
		return
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.App.AllowedOrigins,
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Load validators
	e.Validator = &CustomValidator{validator: validator.New()}

	// API Group
	gApi := e.Group("/api")

	// Websocket (realtime service data)
	gApi.GET("/service/ws", controller.OpenWebsocketConnection)

	// Service CRUD
	gApi.GET("/service/:id", controller.GetServiceWithId)
	gApi.POST("/service", controller.CreateService)
	gApi.PUT("/service", controller.UpdateService)
	gApi.DELETE("/service/:id", controller.DeleteServiceWithId)

	// start background jobs to constantly check
	backgroundCtx := context.Background()
	go tasks.InitBackgroundServiceRefresh(backgroundCtx, time.Duration(config.App.RefreshTimeMs))

	e.Logger.Info(e.Start(":5555"))
	log.Printf("Cleaning up remaining connections")
	backgroundCtx.Done()
}