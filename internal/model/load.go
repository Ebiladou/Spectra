package model

import "time"

type Population struct {
	Size int
}

type LoadProfileType string

const (
	ConstantLoad LoadProfileType = "CONSTANT"
	RampUp       LoadProfileType = "RAMP_UP"
	RampDown     LoadProfileType = "RAMP_DOWN"
)

type LoadProfile struct {
	Type     LoadProfileType
	Duration time.Duration
}
