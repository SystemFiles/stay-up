// API == CONTROLLER
package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/systemfiles/stay-up/api/daos"
	"github.com/systemfiles/stay-up/api/models"
	"github.com/systemfiles/stay-up/api/types"
	"github.com/systemfiles/stay-up/api/util"
)

func CreateService(c echo.Context) error {
	data := new(daos.ServiceCreate)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind service create request. Error: " + err.Error())
	}
	if err := c.Validate(data); err != nil {
		log.Printf("error validating data for create service request")
		return err
	}

	refreshTime := data.RefreshTimeMs
	if refreshTime == 0 {
		refreshTime = 10000
	}

	// Create model from request data
	svc := models.Service{
		Name: data.Name,
		Host: data.Host,
		Port: data.Port,
		Protocol: util.GetProtocolFromString(data.Protocol),
		CurrentStatus: types.UP,
		TimeoutMs: data.TimeoutMs,
		RefreshTimeMs: refreshTime,
	}

	// Open DB connection
	db, err := util.GetDBInstance()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not reach database")
	}

	// Create model in DB
	result := db.Create(&svc)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create the service in target database. " + result.Error.Error())
	}

	return c.JSONPretty(http.StatusCreated, svc, "  ")
}

func GetServiceWithId(c echo.Context) error {
	return nil
}

func UpdateService(c echo.Context) error {
	data := new(daos.ServiceUpdate)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind service update request. Error: " + err.Error())
	}
	if err := c.Validate(data); err != nil {
		log.Printf("error validating data for create service request")
		return err
	}

	// open database connection
	db, err := util.GetDBInstance()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not reach database")
	}

	// perform update to service model
	var svc models.Service
	db.First(&svc, data.ID)
	if svc.ID != data.ID {
		return echo.NewHTTPError(http.StatusNotFound, "Could not find a service with ID, " + fmt.Sprint(data.ID))
	}
	result := db.Model(&svc).Update(data.Attribute, data.NewValue)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Could not perform the update. Reason " + result.Error.Error())
	}

	return c.JSONPretty(http.StatusOK, svc, "  ")
}

func DeleteServiceWithId(c echo.Context) error {
	return nil
}