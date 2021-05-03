package model

import "time"

type SignalKey struct {
	SignalId    string    `json:"signalId" db:"signal_id, omitempty"`
	SignalKey   string    `json:"signalKey" db:"signal_key, omitempty"`
	CreatedAt   time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt   time.Time `json:"-" db:"updated_at, omitempty"`
}

type ScheduleTesting struct {
	ScheduleId    string    `json:"scheduleId" db:"schedule_id, omitempty"`
	Every   string    `json:"every" db:"every, omitempty"`
	Day   string    `json:"day" db:"day, omitempty"`
	AtTime   string    `json:"atTime" db:"at_time, omitempty"`
	SignalKey   string    `json:"signalKey,omitempty" db:"signal_key, omitempty"`
	JobName   string    `json:"jobName,omitempty" db:"job_name, omitempty"`
	CreatedAt   time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt   time.Time `json:"-" db:"updated_at, omitempty"`
}