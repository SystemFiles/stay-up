package provider

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/systemfiles/stay-up/api/daos"
	"github.com/systemfiles/stay-up/api/models"
	"github.com/systemfiles/stay-up/api/types"
	"github.com/systemfiles/stay-up/api/util"
	rclient "github.com/systemfiles/stay-up/api/util/redis"
)

type ServiceProviderError struct{
	Message string
}
func (e *ServiceProviderError) Error() string {
	return fmt.Sprintf("Service Provider Error: %s", e.Message)
}

func CreateService(name, description, host, protocol string, port, timeout int64) (models.Service, error) {
	svcId := uuid.NewString()

	// Create model from request data
	svc := models.Service{
		ID: svcId,
		Name: name,
		Description: description,
		Host: host,
		Port: port,
		Protocol: util.GetProtocolFromString(protocol),
		CurrentStatus: types.UP,
		TimeoutMs: timeout,
		LastDown: time.Now(),
		UptimeSeconds: 0,
		LatencyMs: 0,
	}

	if err := rclient.Set(svcId, &svc); err != nil {
		return models.Service{}, &ServiceProviderError{Message: fmt.Sprintf("Failed to create new service in rdb. Reason: %s", err)}
	}

	return svc, nil
}

func GetServiceById(id string) (models.Service, error) {
	var svc models.Service
	if err := rclient.Get(id, &svc); err != nil {
		return models.Service{}, &ServiceProviderError{Message: fmt.Sprintf("Could not find service with ID, %s", id)}
	}

	return svc, nil
}

func StreamServiceData(ws *websocket.Conn, timeout time.Duration, lastResponse time.Time, errChan chan error) {
	var services []models.Service
	var lastService []models.Service

	for {
		err := ws.WriteMessage(websocket.PingMessage, []byte("keepalive"))
		if err != nil {
			log.Println("Client Not Reachable - Closed connection")
			ws.Close() // Cannot reach client (close the connection)
			return
		}
		if err := GetAllServices(&services); err != nil {
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

func UpdateServiceWithId(newServiceData *daos.ServiceUpdate) (*models.Service, error) {
	old, err := GetServiceById(newServiceData.ID)
	if err != nil {
		return &models.Service{}, err
	}

	svc := &models.Service{
		ID: old.ID,
		Name: newServiceData.Name,
		Description: newServiceData.Description,
		Host: newServiceData.Host,
		Port: newServiceData.Port,
		Protocol: util.GetProtocolFromString(newServiceData.Protocol),
		TimeoutMs: newServiceData.TimeoutMs,
		LastDown: old.LastDown,
		LatencyMs: old.LatencyMs,
		CurrentStatus: old.CurrentStatus,
	}

	return svc, rclient.Set(old.ID, svc)
}

func DeleteServiceWithId(id string) error {
	return rclient.Delete(id)
}

func GetAllServices(dest *[]models.Service) error {
	if err := rclient.GetAll(dest); err != nil {
		return &ServiceProviderError{Message: fmt.Sprintf("Could not get list of services from redis-server. Reason %s", err)}
	}
	return nil
}