package model

import (
	"time"

	"github.com/google/uuid"
)

type Protocol string

const (
	HTTP      Protocol = "HTTP"
	WebSocket Protocol = "WEBSOCKET"
)

type Pause struct {
	Min time.Duration
	Max time.Duration
}

type Scenario struct {
	ID       uuid.UUID
	Name     string
	Protocol Protocol
	UserType string
	Steps    []*Step
}
