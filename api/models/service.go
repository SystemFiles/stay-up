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
	ID uint64 `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Description string `gorm:"not null"`
	Host string `gorm:"not null"`
	Port int64 `gorm:"not null"`
	Protocol types.ServiceProtocol `gorm:"not null"`
	CurrentStatus types.ServiceStatus `gorm:"not null"`
	TimeoutMs int64 `gorm:"not null"`
	LastDown time.Time
	UptimeSeconds int64 `gorm:"check:uptime_seconds >= 0"`
	LatencyMs int64 `gorm:"check:latency_ms >= 0"`
}

func (s Service) Equal(o Service) bool {
	switch {
	case s.ID != o.ID:
		return false
	case s.Name != o.Name:
		return false
	case s.Host != o.Host:
		return false
	case s.Port != o.Port:
		return false
	case s.Description != o.Description:
		return false
	case !reflect.DeepEqual(s.Protocol, o.Protocol):
		return false
	case !reflect.DeepEqual(s.CurrentStatus, o.CurrentStatus):
		return false
	case s.TimeoutMs != o.TimeoutMs:
		return false
	case s.LastDown != o.LastDown:
		return false
	case s.UptimeSeconds != o.UptimeSeconds:
		return false
	case s.LatencyMs != o.LatencyMs:
		return false
	default:
		return true
	}
}

func (s *Service) CheckAndUpdateStatus() error {
	timeout := time.Duration(s.TimeoutMs) * time.Millisecond
	startTime := time.Now()
	_, err := net.DialTimeout(s.Protocol.String(), fmt.Sprintf("%s:%s", s.Host, fmt.Sprint(s.Port)), timeout)
	execTime := time.Since(startTime).Milliseconds()

	// check status for service
	if err != nil {
		s.CurrentStatus = types.DOWN
		s.LastDown = time.Now()
	} else if execTime > time.Duration(100 * time.Millisecond).Milliseconds() {
		s.CurrentStatus = types.SLOW
	} else {
		s.CurrentStatus = types.UP
	}

	// update latency time
	s.LatencyMs = execTime

	// update uptime
	s.UptimeSeconds = int64(time.Since(s.LastDown).Seconds())

	return nil
}