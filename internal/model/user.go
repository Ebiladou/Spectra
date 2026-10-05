package model

import "github.com/google/uuid"

type VirtualUser struct {
	ID        uuid.UUID
	UserType  string
	Scenario  *Scenario
	Variables map[string]string
}
