package automate_plugin

import (
	"github.com/google/uuid"
	"github.com/tebeka/selenium"
	"selenium-check-awingu/helper"
	"selenium-check-awingu/helper/restclient"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/testing"
	"selenium-check-awingu/repository"
	"strconv"
	"time"
)

type RestyTelegram struct {
	RestClient restclient.RestClient
}

func PluginScreenShotError(conctRemote selenium.WebDriver, contentTesting testing.YamlTesting, userSignIn model.JobsUser,
	testingRepo repository.TestingRepo, plusInfoActionTesting testing.PlusInfoActionTesting,
	restClient restclient.RestClient) error {
	pIAT := plusInfoActionTesting

	shot, err := conctRemote.Screenshot()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	ramdomId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	nameImage := contentTesting.NameTest + "-" + userSignIn.Username + "-" + time.Now().Format("20060102150405") + "-" + ramdomId.String()

	err = helper.HelpSaveImage(shot, nameImage)
	if err != nil {
		log.Error(err.Error())
		return err
	} else {
		log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) + " Đã lưu ảnh thành công " + nameImage)
	}

	stringImage, err := helper.Base64ImageToString(nameImage)
	if err != nil {
		log.Error(err.Error())
		return err
	} else {
		log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) + " Đã base64 " + nameImage + " sang string.")
	}
	pIAT.WebDriver = "Screenshot"
	pIAT.Action = "Error"
	pIAT.Data = stringImage
	err = RecordActionToDB(testingRepo,
		contentTesting, userSignIn,
		pIAT)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	messageAlert := contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) +
		" " + pIAT.DescriptionStep + " Hoàn thành ghi nhận error"
	if pIAT.AlertTelegram != "no" {
		log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) +
			" " + pIAT.DescriptionStep + " Gửi TeleGram ghi nhận error")
		err = restClient.SendMessageToGroupTelegram(messageAlert, pIAT.TelegramInfo)
	}
	//////////END log action
	log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) +
		" " + pIAT.DescriptionStep + " Hoàn thành ghi nhận error")
	return nil
}
