package models

import (
	"fmt"
	"net"
	"reflect"
	"time"

	"github.com/systemfiles/stay-up/api/types"
)

type IService interface {
	CheckStatus() error
}

type Service struct {
	ID string
	Name string
	Description string
	Host string
	Port int64
	Protocol types.ServiceProtocol
	CurrentStatus types.ServiceStatus
	TimeoutMs int64
	LastDown time.Time
	UptimeSeconds int64
	LatencyMs int64
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
	
	// perform uptime test with TCP / UDP connect
	startTime := time.Now()
	_, err := net.DialTimeout(s.Protocol.String(), fmt.Sprintf("%s:%s", s.Host, fmt.Sprint(s.Port)), timeout)
	execTime := time.Since(startTime).Milliseconds()

	// update status with conditions
	if err != nil {
		s.CurrentStatus = types.DOWN
		s.LatencyMs = 0
		s.LastDown = time.Now()
	} else {
		if execTime > time.Duration(100 * time.Millisecond).Milliseconds() {
			s.CurrentStatus = types.SLOW
			s.LatencyMs = execTime
		} else {
			s.CurrentStatus = types.UP
			s.LatencyMs = execTime
		}
	}

	// update uptime
	s.UptimeSeconds = int64(time.Since(s.LastDown).Seconds())

	return nil
}