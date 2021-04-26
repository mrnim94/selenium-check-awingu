package handler

import (
	"context"
	"selenium-check-awingu/banana"
	"selenium-check-awingu/helper/automate"
	"selenium-check-awingu/log"
	"selenium-check-awingu/repository"
	"time"
)

type AutoTestingHandler struct {
	UserTesting repository.UserTesting
	Automate    automate.Automate
}

func (a *AutoTestingHandler) HandlerTestingConcurrency() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	users, err := a.UserTesting.SelectAllUsersBooking(ctx)
	if err != nil {
		if err == banana.UserBookingNotFound {
			log.Error(err.Error())
			return banana.UserBookingNotFound
		}
		log.Error(err.Error())
		return err
	}

	for i, user := range users {
		log.Printf("Thông tin users lấy được", users[i])
		go a.Automate.TestingConcurrencyAwingu(user)
	}

	return nil
}
