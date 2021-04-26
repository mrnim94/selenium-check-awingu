package model

import "time"

type Testing struct {
	UserId           string    `json:"-" db:"user_id, omitempty"`
	Username         string    `json:"userName,omitempty" db:"username, omitempty"`
	Password         string    `json:"password,omitempty" db:"password, omitempty"`
	CreatedTesting   time.Time `json:"createdTesting,omitempty" db:"created_testing, omitempty"`
	CompletedTesting time.Time `json:"completedTesting,omitempty" db:"completed_testing, omitempty"`
}
