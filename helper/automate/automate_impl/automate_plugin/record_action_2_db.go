package automate_plugin

import (
	"context"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/testing"
	"selenium-check-awingu/repository"
	"time"
)

func RecordActionToDB(testingRepo repository.TestingRepo,
	contentTesting testing.YamlTesting, userSignIn model.JobsUser,
	plusInfoActionTesting testing.PlusInfoActionTesting) error {
	uAT := testing.UserActionTesting{
		JobId:       plusInfoActionTesting.JobID,
		TimeId:      time.Now().UnixNano(),
		TestId:      plusInfoActionTesting.TestId,
		NameTest:    contentTesting.NameTest,
		Version:     contentTesting.Version,
		Browser:     contentTesting.Browser,
		Page:        plusInfoActionTesting.Page,
		Agent:       userSignIn.Username,
		WebDriver:   plusInfoActionTesting.WebDriver,
		Description: plusInfoActionTesting.DescriptionStep,
		Action:      plusInfoActionTesting.Action,
		Data:        plusInfoActionTesting.Data,
		Data1:       plusInfoActionTesting.Data1,
	}

	context := context.Background()
	_, err := testingRepo.SaveActionOfUser(context, uAT)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
