package model

import (
	"time"

	"github.com/google/uuid"
)

type Pause struct {
	Min time.Duration
	Max time.Duration
}

type Step struct {
	ID     uuid.UUID
	Action string
	Method string
	URL    string
	Data   string
	Pause  *Pause
}
