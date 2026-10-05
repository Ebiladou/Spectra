package model

import (
	"github.com/google/uuid"
)

type Metric string

const (
	RequestCount Metric = "REQUEST_COUNT"
	SuccessCount Metric = "SUCCESS_COUNT"
	ErrorCount   Metric = "ERROR_COUNT"
	ErrorRate    Metric = "ERROR_RATE"
	Duration     Metric = "DURATION"
	Throughput   Metric = "THROUGHPUT"
)

type Percentile string

const (
	P50 Percentile = "P50"
	P90 Percentile = "P90"
	P95 Percentile = "P95"
	P99 Percentile = "P99"
	Max Percentile = "MAX"
)

type ThresholdOperator string

const (
	LessThan           ThresholdOperator = "<"
	LessThanOrEqual    ThresholdOperator = "<="
	GreaterThan        ThresholdOperator = ">"
	GreaterThanOrEqual ThresholdOperator = ">="
	Equal              ThresholdOperator = "=="
)

type MetricUnit string

const (
	Count             MetricUnit = "COUNT"
	Percentage        MetricUnit = "PERCENTAGE"
	DurationMs        MetricUnit = "DURATION_MS"
	RequestsPerSecond MetricUnit = "REQUESTS_PER_SECOND"
)

type MetricValue struct {
	Number float64
	Unit   MetricUnit
}

type Threshold struct {
	Metric     Metric
	Percentile *Percentile
	Operator   ThresholdOperator
	Value      MetricValue
}

type Test struct {
	ID          uuid.UUID
	Name        string
	Population  Population
	LoadProfile LoadProfile
	Scenarios   []*ScenarioAssignment
	Thresholds  []Threshold
}
