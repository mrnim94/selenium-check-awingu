package model

import "time"

type JobsGithub struct {
	GithubId    string    `json:"-" db:"github_id, omitempty"`
	AccessToken string    `json:"-" db:"access_token, omitempty"`
	Owner       string    `json:"owner,omitempty" db:"owner, omitempty"`
	Repo        string    `json:"repo,omitempty" db:"repo, omitempty"`
	Path        string    `json:"path,omitempty" db:"path, omitempty"`
	JobID       string    `json:"jobID,omitempty" db:"job_id, omitempty"`
	CreatedAt   time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt   time.Time `json:"-" db:"updated_at, omitempty"`
}
