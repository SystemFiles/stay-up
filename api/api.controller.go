// API == CONTROLLER
package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/systemfiles/stay-up/api/daos"
	"github.com/systemfiles/stay-up/api/models"
	"github.com/systemfiles/stay-up/api/types"
	"github.com/systemfiles/stay-up/api/util"
)

func CreateService(c echo.Context) error {
	data := new(daos.Create)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind service create request. Error: " + err.Error())
	}
	if err := c.Validate(data); err != nil {
		log.Printf("ERROR VALIDATING")
		return err
	}

	// Create model from request data
	svc := models.Service{
		Name: data.Name,
		Host: data.Host,
		Port: data.Port,
		Protocol: util.GetProtocolFromString(data.Protocol),
		CurrentStatus: types.UP,
		TimeoutMs: data.TimeoutMs,
		RefreshTimeMs: data.RefreshTimeMs,
	}

	// Open DB connection
	db, err := util.GetDBInstance()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not read database")
	}

	// Create model in DB
	result := db.Statement.Create(&svc)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create the service in target database")
	}

	return c.JSONPretty(http.StatusCreated, svc, "  ")
}

func GetStatus(c echo.Context) error {
	svc := models.Service{
		Name: "Bitwarden",
		Host: "bitwarden.sykesdev.ca",
		Port: 443,
		Protocol: types.PROTO_TCP,
		CurrentStatus: types.UP,
		TimeoutMs: 5096,
		RefreshTimeMs: 200,
	}

		err := svc.CheckStatus()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, svc)
}