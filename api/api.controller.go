// API == CONTROLLER
package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/systemfiles/stay-up/api/daos"
	"github.com/systemfiles/stay-up/api/provider"
)

func CreateService(c echo.Context) error {
	data := new(daos.ServiceCreate)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind service create request. Error: " + err.Error())
	}
	if err := c.Validate(data); err != nil {
		return err
	}

	// create service in database
	svc, err := provider.CreateService(data.Name, data.Host, data.Protocol, data.Port, data.TimeoutMs)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSONPretty(http.StatusCreated, svc, "  ")
}

func GetServiceWithId(c echo.Context) error {
	serviceID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Cannot process ID provided")
	}

	// get service from database
	svc, err := provider.GetServiceById(serviceID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
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

	// perform update using provider
	svc, err := provider.UpdateServiceWithId(data.ID, data.Attribute, data.NewValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSONPretty(http.StatusOK, svc, "  ")
}

func DeleteServiceWithId(c echo.Context) error {
	serviceID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Cannot process ID provided")
	}

	// perform delete using provider
	if err := provider.DeleteServiceWithId(serviceID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSONPretty(http.StatusOK, daos.DeleteServiceResponse{Message: "Deleted service with ID " + fmt.Sprint(serviceID)}, "  ")
}