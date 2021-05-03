package model

import "time"

type JobsTesting struct {
	JobId     string    `json:"jobID,omitempty" db:"job_id, omitempty"`
	JobName   string    `json:"jobName,omitempty" db:"job_name, omitempty"`
	Status    int    `json:"status,omitempty" db:"status, omitempty"`
	AlertTelegram   string    `json:"alertTelegram,omitempty" db:"alert_telegram, omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at, omitempty"`
}
