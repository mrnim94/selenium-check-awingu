package automate_plugin

import (
	"github.com/tebeka/selenium"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/testing"
	"selenium-check-awingu/repository"
	"strconv"
	"time"
)

func SetImplicitWaitTimeout(conctRemote selenium.WebDriver, contentTesting testing.YamlTesting, userSignIn model.JobsUser,
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
	log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) + " " + pIAT.DescriptionStep+" Begin")
	timeOut, _ := strconv.Atoi(pIAT.Timeout)
	conctRemote.SetImplicitWaitTimeout(time.Second * time.Duration(timeOut))

	pIAT.Action = "End"
	err = RecordActionToDB(testingRepo,
		contentTesting, userSignIn,
		pIAT)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) + " " + pIAT.DescriptionStep+" End")
	//////////END log action
	//End case "SetImplicitWaitTimeout":
	return nil
}
