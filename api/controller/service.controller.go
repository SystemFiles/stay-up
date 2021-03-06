package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/systemfiles/stay-up/api/daos"
	"github.com/systemfiles/stay-up/api/provider"
)

var wsUpgrader = websocket.Upgrader{}

func OpenWebsocketConnection(c echo.Context) error {
	wsTimeout := time.Duration(5 * time.Second)
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {return true}

	ws, err := wsUpgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, fmt.Sprintf("Could not establish a reliable connection to the websocket. Reason: %s", err.Error()))
	}
	defer ws.Close()
	
	log.Println("Connected")

	// Setup socket to stream the service data to the connected client
	errChan := make(chan error)
	lastResponse := time.Now()
	ws.SetPongHandler(func(msg string) error {
		log.Printf("Got Ping Message: %s\n", msg)
		lastResponse = time.Now()
		return nil
	})

	go provider.StreamServiceData(ws, wsTimeout, lastResponse, errChan)
	
	if err := <- errChan; err != nil {
		log.Println("ERROR OCCURRED: Closing websocket connection")
		ws.Close()

		return err
	}

	return nil
}

func CreateService(c echo.Context) error {
	data := new(daos.ServiceCreate)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind service create request. Error: " + err.Error())
	}
	if err := c.Validate(data); err != nil {
		return err
	}

	// create service in database
	svc, err := provider.CreateService(data.Name, data.Description, data.Host, data.Protocol, data.Port, data.TimeoutMs)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSONPretty(http.StatusCreated, svc, "  ")
}

func GetServiceWithId(c echo.Context) error {
	serviceID := c.Param("id")
	
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
	svc, err := provider.UpdateServiceWithId(data)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSONPretty(http.StatusOK, svc, "  ")
}

func DeleteServiceWithId(c echo.Context) error {
	serviceID := c.Param("id")
	
	// perform delete using provider
	if err := provider.DeleteServiceWithId(serviceID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSONPretty(http.StatusOK, daos.DeleteServiceResponse{Message: "Deleted service with ID " + fmt.Sprint(serviceID)}, "  ")
}