package model

import (
	"time"

	"github.com/google/uuid"
)

type Sample struct {
	Timestamp     time.Time
	ScenarioID    uuid.UUID
	StepID        uuid.UUID
	VirtualUserID uuid.UUID
	Duration      time.Duration
	Success       bool
	StatusCode    int
	Error         error
}

type Metrics struct {
	TotalRequests int64
	Successful    int64
	Failed        int64
	TotalDuration time.Duration
}

type TestResult struct {
	TestID     uuid.UUID
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   time.Duration
	Metrics    Metrics
	Thresholds []ThresholdResult
}

type ThresholdResult struct {
	Threshold Threshold
	Passed    bool
	Actual    float64
}
