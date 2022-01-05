// API == CONTROLLER
package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/systemfiles/stay-up/api/daos"
	"github.com/systemfiles/stay-up/api/models"
	"github.com/systemfiles/stay-up/api/provider"
)

var wsUpgrader = websocket.Upgrader{}

func streamServiceData(ws *websocket.Conn, timeout time.Duration, lastResponse time.Time, errChan chan error) {
	var services []models.Service
	var lastService []models.Service

	for {
		err := ws.WriteMessage(websocket.PingMessage, []byte("keepalive"))
		if err != nil {
			log.Println("Websocket connection could not stay alive...")
			errChan <- echo.NewHTTPError(http.StatusInternalServerError, "Websocket connection could not stay alive...")
		}
		if err := provider.GetAllServices(&services); err != nil {
			log.Println("Failed to get service from database")
			errChan <- echo.NewHTTPError(http.StatusInternalServerError, "Failed to get services from data source")
		}
		if !cmp.Equal(services, lastService) {
			// If data has changed from the first sent data then send updated data
			if err := ws.WriteJSON(services); err != nil {
				log.Println(fmt.Sprintf("Websocket write failed ... %s", err.Error()))
				errChan <- echo.NewHTTPError(http.StatusInternalServerError, "Websocket write failed ... ")
			}
			lastResponse = time.Now()
			lastService = services
		} else {
			lastResponse = time.Now()
		}
		// wait before next iteration
		time.Sleep(time.Duration(timeout / 2))
		if (time.Since(lastResponse) > timeout) {
			log.Println("Closed connection")
			ws.Close()
			return
		}
	}
}

func OpenWebsocketConnection(c echo.Context) error {
	wsTimeout := time.Duration(2000 * time.Millisecond)
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {return true}

	ws, err := wsUpgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, fmt.Sprintf("Could not establish a reliable connection to the websocket. Reason: %s", err.Error()))
	}
	defer ws.Close()
	
	log.Println("Connected")

	errChan := make(chan error)
	lastResponse := time.Now()
	ws.SetPongHandler(func(msg string) error {
		log.Printf("Got Ping Message: %s\n", msg)
		lastResponse = time.Now()
		return nil
	})

	go streamServiceData(ws, wsTimeout, lastResponse, errChan)
	return <- errChan
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