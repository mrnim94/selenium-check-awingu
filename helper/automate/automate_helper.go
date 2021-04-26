package automate

import (
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/testing"
	"selenium-check-awingu/repository"
)

type Automate interface {
	TestingConcurrencyAwingu(testing model.Testing) error
	RobotAutoImpl(contentTesting testing.YamlTesting, userSignIn model.JobsUser, testingRepo repository.TestingRepo, jodID string) error
}
