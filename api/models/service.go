package models

import (
	"fmt"
	"net"
	"reflect"
	"time"

	"github.com/systemfiles/stay-up/api/types"
	"gorm.io/gorm"
)

type IService interface {
	CheckStatus() error
	ModifyAttribute(key string, value interface{}) error
}

type Service struct {
	gorm.Model
	Name string
	Host string
	Port int64
	Protocol types.ServiceProtocol
	CurrentStatus types.ServiceStatus
	TimeoutMs int64
	RefreshTimeMs int64
}

func (s *Service) CheckStatus() error {
	timeout := time.Duration(s.TimeoutMs) * time.Millisecond
	startTime := time.Now()
	conn, err := net.DialTimeout(s.Protocol.String(), fmt.Sprintf("%s:%s", s.Host, fmt.Sprint(s.Port)), timeout)
	if err != nil {
		s.CurrentStatus = types.DOWN
		return nil
	}
	execTime := time.Since(startTime).Milliseconds()
	fmt.Printf("Took %d ms to reach %s\n", execTime, s.Name)
	
	if execTime > time.Duration(50 * time.Millisecond).Milliseconds() {
		s.CurrentStatus = types.SLOW
		return nil
	}

	conn.Close()
	return nil
}

func (s *Service) ModifyAttribute(key string, value interface{}) error {
	rvalue := reflect.ValueOf(value)
	serviceFields := reflect.ValueOf(s).Elem()

	fieldLookup := serviceFields.FieldByName(key)
	if !fieldLookup.IsValid() {
		return fmt.Errorf("%s is not a valid field name: ", key)
	}

	if !fieldLookup.CanSet() && key != "CurrentStatus" {
		return fmt.Errorf("cannot set value for %s", key)
	}

	if fieldLookup.Kind() != rvalue.Kind() {
		return fmt.Errorf("the value specified for %s is not the same", key)
	}

	fieldLookup.Set(rvalue)
	return nil
}