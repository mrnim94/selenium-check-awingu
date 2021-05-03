package restclient_impl

import (
	"crypto/tls"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model/req"
	"selenium-check-awingu/banana"
)

func (r *Resty) RunTesting(jobName string) error{
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	body := req.RequestTesting{
		JobName: jobName,
	}
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		//SetHeader("X-Auth-Token", token).
		SetBody(body).
		SetResult(&AuthSuccess{}). // or SetResult(AuthSuccess{}).
		SetError(&AuthError{}). // or SetError(AuthError{}).
		Post(r.Url+"/tester/run-testing")
	if err != nil {
		log.Error(err.Error())
		return  banana.RunTestingScheduleError
	}

	return  nil
}