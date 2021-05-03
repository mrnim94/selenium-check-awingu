package automate

import (
	"selenium-check-awingu/helper/restclient"
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/alert"
	"selenium-check-awingu/model/testing"
	"selenium-check-awingu/repository"
)

type Automate interface {
	TestingConcurrencyAwingu(testing model.Testing) error
	RobotAutoImpl(contentTesting testing.YamlTesting, userSignIn model.JobsUser,
					testingRepo repository.TestingRepo, job model.JobsTesting, teleInfo alert.TelegramInfo,
					restClient restclient.RestClient) error
}
