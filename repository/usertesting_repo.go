package repository

import (
	"context"
	"selenium-check-awingu/model"
)

type UserTesting interface {
	SelectAllUsersBooking(context context.Context) ([]model.Testing, error)
}
