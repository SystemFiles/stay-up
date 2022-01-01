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
}

type Service struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Host string `gorm:"not null"`
	Port int64 `gorm:"not null"`
	Protocol types.ServiceProtocol `gorm:"not null"`
	CurrentStatus types.ServiceStatus `gorm:"not null"`
	TimeoutMs int64 `gorm:"not null"`
	RefreshTimeMs int64 `gorm:"not null"`
}

func (s *Service) CheckStatus() error {
	timeout := time.Duration(s.TimeoutMs) * time.Millisecond
	startTime := time.Now()
	conn, err := net.DialTimeout(s.Protocol.String(), fmt.Sprintf("%s:%s", s.Host, fmt.Sprint(s.Port)), timeout)
	if err != nil {
		s.CurrentStatus = types.DOWN
	}
	execTime := time.Since(startTime).Milliseconds()
	
	if execTime > time.Duration(50 * time.Millisecond).Milliseconds() {
		s.CurrentStatus = types.SLOW
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