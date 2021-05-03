package restclient

import (
	"selenium-check-awingu/model/alert"
)

type RestClient interface {
	RunTesting(jobName string) error

	//Call API telegram
	SendMessageToGroupTelegram(messageAlert string, teleInfo alert.TelegramInfo) error
}