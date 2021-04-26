package automate_plugin

import (
	"github.com/tebeka/selenium"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/testing"
	"selenium-check-awingu/repository"
	"strconv"
)

func AccessNewUrl(conctRemote selenium.WebDriver, contentTesting testing.YamlTesting, userSignIn model.JobsUser,
	testingRepo repository.TestingRepo, plusInfoActionTesting testing.PlusInfoActionTesting) error {
	pIAT := plusInfoActionTesting
	pIAT.Action = "Begin"

	err := RecordActionToDB(testingRepo,
		contentTesting, userSignIn,
		pIAT)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	//////////END log action
	log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) + " " + pIAT.DescriptionStep)

	err = conctRemote.Get(pIAT.NewUrl)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	pIAT.Action = "End"

	err = RecordActionToDB(testingRepo,
		contentTesting, userSignIn,
		pIAT)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	//////////END log action
	log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) +
		" " + pIAT.DescriptionStep + " Hoàn thành")
	return nil
}
