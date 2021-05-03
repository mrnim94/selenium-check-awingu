package report

import "time"

type ReportJobsTesting struct {
	JobId     string    `json:"jobID,omitempty" db:"job_id, omitempty"`
	JobName   string    `json:"jobName,omitempty" db:"job_name, omitempty"`
	Status    string    `json:"status,omitempty" db:"status, omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at, omitempty"`
	Owner     string    `json:"owner,omitempty" db:"owner, omitempty"`
	Repo      string    `json:"repo,omitempty" db:"repo, omitempty"`
	Path      string    `json:"path,omitempty" db:"path, omitempty"`
	AlertTelegram   string    `json:"alertTelegram,omitempty" db:"alert_telegram, omitempty"`
}
