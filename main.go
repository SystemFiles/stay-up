package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/systemfiles/stay-up/api"
	"github.com/systemfiles/stay-up/api/config"
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

	// Load validators
	e.Validator = &CustomValidator{validator: validator.New()}

	// Websocker (realtime data)
	e.GET("/", func(c echo.Context) error {
		// Open websocker
		return echo.NewHTTPError(http.StatusNotImplemented, "This endpoint is not implemented")
	})

	// API Group
	gApi := e.Group("/api")

	// Service CRUD
	gApi.GET("/service/:id", api.GetServiceWithId)
	gApi.POST("/service", api.CreateService)
	gApi.PUT("/service", api.UpdateService)
	gApi.DELETE("/service/:id", api.DeleteServiceWithId)

	// start background jobs to constantly check
	backgroundCtx := context.Background()
	go tasks.InitBackgroundServiceRefresh(backgroundCtx)

	e.Logger.Info(e.Start(":5555"))
	log.Printf("Cleaning up remaining connections")
	backgroundCtx.Done()
}