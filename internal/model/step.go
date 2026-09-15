package model

import "github.com/google/uuid"

type Step struct {
	ID     uuid.UUID
	Action string
	Method string
	URL    string
	Data   string
	Pause  *Pause
}
