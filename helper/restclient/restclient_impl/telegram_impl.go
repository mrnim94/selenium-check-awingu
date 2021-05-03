package restclient_impl

import (
	"crypto/tls"
	"selenium-check-awingu/banana"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model/alert"
	"selenium-check-awingu/model/req"
)

func (r *Resty) SendMessageToGroupTelegram(messageAlert string, teleInfo alert.TelegramInfo) error{

	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	body := req.RequestSendMessageToGroup{
		ChatId:              teleInfo.ChatId,
		Text:                messageAlert,
		DisableNotification: false,
	}
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		//SetHeader("X-Auth-Token", token).
		SetBody(body).
		SetResult(&AuthSuccess{}). // or SetResult(AuthSuccess{}).
		SetError(&AuthError{}). // or SetError(AuthError{}).
		Post(r.Url+"/bot"+teleInfo.TelegramToken+"/sendMessage")
	if err != nil {
		log.Error(err.Error())
		return  banana.RunTestingScheduleError
	}

	return  nil
}
