package model

import "github.com/google/uuid"

type ExecutionContext struct {
	ScenarioID    uuid.UUID
	VirtualUserID uuid.UUID
	Variables     map[string]string
}
