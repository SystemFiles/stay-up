package models

import (
	"fmt"
	"net"
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
	Host string `gorm:"not null"`
	Port int64 `gorm:"not null"`
	Protocol types.ServiceProtocol `gorm:"not null"`
	CurrentStatus types.ServiceStatus `gorm:"not null"`
	TimeoutMs int64 `gorm:"not null"`
}

func (s *Service) CheckAndUpdateStatus() (*net.Conn, error) {
	timeout := time.Duration(s.TimeoutMs) * time.Millisecond
	startTime := time.Now()
	conn, err := net.DialTimeout(s.Protocol.String(), fmt.Sprintf("%s:%s", s.Host, fmt.Sprint(s.Port)), timeout)
	execTime := time.Since(startTime).Milliseconds()
	if err != nil {
		s.CurrentStatus = types.DOWN
	} else if execTime > time.Duration(100 * time.Millisecond).Milliseconds() {
		s.CurrentStatus = types.SLOW
	} else {
		s.CurrentStatus = types.UP
	}

	return &conn, nil
}