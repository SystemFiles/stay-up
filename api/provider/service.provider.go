package provider

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/systemfiles/stay-up/api/models"
	"github.com/systemfiles/stay-up/api/types"
	"github.com/systemfiles/stay-up/api/util"
	"gorm.io/gorm/clause"
)

type ServiceProviderError struct{
	Message string
}
func (e *ServiceProviderError) Error() string {
	return fmt.Sprintf("Service Provider Error: %s", e.Message)
}

func CreateService(name, description, host, protocol string, port, timeout int64) (models.Service, error) {
	// Create model from request data
	svc := models.Service{
		Name: name,
		Description: description,
		Host: host,
		Port: port,
		Protocol: util.GetProtocolFromString(protocol),
		CurrentStatus: types.UP,
		TimeoutMs: timeout,
	}

	// Open DB connection
	db, err := util.GetDBInstance()
	if err != nil {
		return models.Service{}, &ServiceProviderError{Message: "Could not get a valid database instance"}
	}

	// Create model in DB
	if err := db.Create(&svc).Error; err != nil {
		return models.Service{}, &ServiceProviderError{Message: fmt.Sprintf("Failed to create the service in target database. Reason: %s", err)}
	}

	return svc, nil
}

func GetServiceById(id uint64) (models.Service, error) {
	// open database connection
	db, err := util.GetDBInstance()
	if err != nil {
		return models.Service{}, &ServiceProviderError{Message: "Could not get a valid database instance"}
	}

	// find service with id
	var svc models.Service
	db.First(&svc, id)
	if svc.ID != id {
		return models.Service{}, &ServiceProviderError{Message: fmt.Sprintf("Could not find service with ID, %d.", id)}
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

func UpdateServiceWithId(id uint64, attr string, val interface{}) (models.Service, error) {
	// open database connection
	db, err := util.GetDBInstance()
	if err != nil {
		return models.Service{}, &ServiceProviderError{Message: "Could not get a valid database instance"}
	}

	// find service model with given primary_key -> id
	var svc models.Service
	db.First(&svc, id)
	if svc.ID != id {
		return models.Service{}, &ServiceProviderError{Message: fmt.Sprintf("Could not find service with ID, %d.", id)}
	}

	// make update
	if err := db.Model(&svc).Update(attr, val).Error; err != nil {
		return models.Service{}, &ServiceProviderError{Message: fmt.Sprintf("Could not perform the update. Reason %s", err.Error())}
	}

	return svc, nil
}

func DeleteServiceWithId(id uint64) error {
	// open database connection
	db, err := util.GetDBInstance()
	if err != nil {
		return &ServiceProviderError{Message: "Could not get a valid database instance"}
	}

	// find service using ID
	var svc models.Service
	db.First(&svc, id)
	if svc.ID != id {
		return &ServiceProviderError{Message: fmt.Sprintf("Could not find service with ID, %d.", id)}
	}

	// delete the service
	if err := db.Delete(&svc).Error; err != nil {
		return &ServiceProviderError{Message: fmt.Sprintf("Failed to delete service from database. Reason: %s", err)}
	}
	
	return nil
}

func GetAllServices(dest *[]models.Service) error {
	// open database connection
	db, err := util.GetDBInstance()
	if err != nil {
		return err
	}

	// get services from database
	if err := db.Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: false}).Find(&dest).Error; err != nil {
		return err
	}

	return nil
}