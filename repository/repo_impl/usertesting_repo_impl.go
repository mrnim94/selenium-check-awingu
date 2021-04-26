package repo_impl

import (
	"context"
	"database/sql"
	"selenium-check-awingu/banana"
	"selenium-check-awingu/db"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	"selenium-check-awingu/repository"
)

type UserBookingRepoImpl struct {
	sql *db.Sql
}

func NewUserTesting(sql *db.Sql) repository.UserTesting {
	return &UserBookingRepoImpl{
		sql: sql,
	}
}

func (u *UserBookingRepoImpl) SelectAllUsersBooking(context context.Context) ([]model.Testing, error) {
	users := []model.Testing{}
	err := u.sql.Db.SelectContext(context, &users,
		`SELECT user_id, username, password
				FROM users_testing`)

	if err != nil {
		if err == sql.ErrNoRows {
			return users, banana.UserBookingNotFound
		}
		log.Error(err.Error())
		return users, err
	}
	return users, nil
}
