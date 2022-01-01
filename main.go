package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/systemfiles/stay-up/api"
	"github.com/systemfiles/stay-up/api/config"
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

	// Client (no group req)
	e.GET("/", func(c echo.Context) error {
		
		return c.HTML(http.StatusOK, "<h1>This is the main page of the website</h1>")
	})

	// API Group
	gApi := e.Group("/api")
	gApi.POST("/service", api.CreateService)

	e.Logger.Fatal(e.Start(":5555"))
}