package model

import (
	"github.com/google/uuid"
)

type Protocol string

const (
	HTTP      Protocol = "HTTP"
	WebSocket Protocol = "WEBSOCKET"
)

type Scenario struct {
	ID       uuid.UUID
	Name     string
	Protocol Protocol
	UserType string
	Steps    []*Step
}

type ScenarioAssignment struct {
	Scenario *Scenario
	Weight   int
}
