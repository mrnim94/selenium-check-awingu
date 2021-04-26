package model

import "time"

type JobsUser struct {
	UserId    string    `json:"-" db:"user_id, omitempty"`
	Username  string    `json:"userName,omitempty" db:"username, omitempty"`
	Password  string    `json:"-" db:"password, omitempty"`
	JobId     string    `json:"jobId,omitempty" db:"job_id, omitempty"`
	Status    int    `json:"status,omitempty" db:"status, omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at, omitempty"`
}
