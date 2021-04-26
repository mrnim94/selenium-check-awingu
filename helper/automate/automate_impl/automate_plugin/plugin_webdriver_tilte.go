package automate_plugin

import (
	"github.com/tebeka/selenium"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/testing"
	"selenium-check-awingu/repository"
	"strconv"
)

/// Funtion lấy tilte
func WebDriverTitle(conctRemote selenium.WebDriver, contentTesting testing.YamlTesting, userSignIn model.JobsUser,
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

	var title string
	if title, err = conctRemote.Title(); err == nil {
		log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) + " bạn đang test trang: " + title)
	} else {
		log.Error(err.Error())
		return err
	}

	pIAT.Action = "End"
	pIAT.Data = title

	err = RecordActionToDB(testingRepo,
		contentTesting, userSignIn,
		pIAT)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	//////////END log action
	return nil
}
