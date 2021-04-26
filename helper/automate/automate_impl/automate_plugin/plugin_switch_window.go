package automate_plugin

import (
	"github.com/tebeka/selenium"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/testing"
	"selenium-check-awingu/repository"
	"strconv"
	"errors"
)

func SwitchWindow(conctRemote selenium.WebDriver, contentTesting testing.YamlTesting, userSignIn model.JobsUser,
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
	nameTabBrowser, err := conctRemote.WindowHandles()
	log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) + " Hiện tại Browser của bạn có: " + strconv.Itoa(len(nameTabBrowser)) + " tab")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	tab, _ := strconv.Atoi(pIAT.Tab)
	if tab >= len(nameTabBrowser) {
		log.Error("Khai báo của bạn chuyển tab của có giá trị: "+pIAT.Tab+ " nhỏ hơn số tab hiện có là "+ strconv.Itoa(len(nameTabBrowser)))
		PluginScreenShotWarning(conctRemote, contentTesting, userSignIn,
			testingRepo, plusInfoActionTesting)
		return errors.New("error: index out of range")
	}
	tabId := nameTabBrowser[tab]
	err = conctRemote.SwitchWindow(tabId)
	if err != nil {
		log.Error(err.Error())
		return err
	} else {
		pIAT.Action = "End"
		err := RecordActionToDB(testingRepo,
			contentTesting, userSignIn,
			pIAT)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		//////////END log action
		log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) + " " + pIAT.DescriptionStep + " hoàn thành")
	}
	return nil
	//End case "SwitchWindow":
}
