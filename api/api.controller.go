// API == CONTROLLER
package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	// Create model from request data
	svc := models.Service{
		Name: data.Name,
		Host: data.Host,
		Port: data.Port,
		Protocol: util.GetProtocolFromString(data.Protocol),
		CurrentStatus: types.UP,
		TimeoutMs: data.TimeoutMs,
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
	serviceID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Cannot process ID provided")
	}

	// open database connection
	db, err := util.GetDBInstance()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not reach database")
	}

	// find service with id
	var svc models.Service
	db.First(&svc, serviceID)
	if svc.ID != serviceID {
		return echo.NewHTTPError(http.StatusNotFound, "Could not find the desired service")
	}

	return c.JSONPretty(http.StatusOK, svc, "  ")
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

	// find service model with given primary_key -> id
	var svc models.Service
	db.First(&svc, data.ID)
	if svc.ID != data.ID {
		return echo.NewHTTPError(http.StatusNotFound, "Could not find a service with ID, " + fmt.Sprint(data.ID))
	}

	// make update
	if err := db.Model(&svc).Update(data.Attribute, data.NewValue).Error; err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Could not perform the update. Reason " + err.Error())
	}

	return c.JSONPretty(http.StatusOK, svc, "  ")
}

func DeleteServiceWithId(c echo.Context) error {
	serviceID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Cannot process ID provided")
	}

	// open database connection
	db, err := util.GetDBInstance()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not reach database")
	}

	// find service using ID
	var svc models.Service
	db.First(&svc, serviceID)
	if svc.ID != serviceID {
		return echo.NewHTTPError(http.StatusNotFound, "Could not find the desired service for deletion")
	}

	// delete the service
	if err := db.Delete(&svc).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete service from database. " + err.Error())
	}

	return c.JSONPretty(http.StatusOK, daos.DeleteServiceResponse{Message: "Deleted service with ID " + fmt.Sprint(serviceID)}, "  ")
}