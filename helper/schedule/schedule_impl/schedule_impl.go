package schedule_impl

import (
	"selenium-check-awingu/helper/restclient"
	"selenium-check-awingu/helper/schedule"
	"selenium-check-awingu/model"
	"selenium-check-awingu/log"
	"strconv"
)

type TestingScheduleImpl struct {
	goCron *schedule.GoCron
}

func NewTestingSchedule(goCron *schedule.GoCron) schedule.TestingSchedule  {
	return &TestingScheduleImpl {
		goCron: goCron,
	}
}


func (ts *TestingScheduleImpl) AddJobAtSpecificTime(scheduleT model.ScheduleTesting, restC restclient.RestClient) error {
	intervalSecond, err := strconv.Atoi(scheduleT.Every)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	job, err := ts.goCron.GR.Every(intervalSecond).Day().At(scheduleT.AtTime).Tag(scheduleT.ScheduleId).Do(func() error {
		err := restC.RunTesting(scheduleT.JobName)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		log.Error(err.Error())
		return err
	}
	ts.goCron.GR.StartAsync()
	log.Info("Schedule có ID " +scheduleT.ScheduleId+" sẽ chạy vào lúc "+job.ScheduledAtTime())
	return nil
}

func (ts *TestingScheduleImpl) RemoveScheduleByTag(scheduleId string) error  {
	err := ts.goCron.GR.RemoveByTag(scheduleId)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info("Thực hiện Remove scheduler có id : "+ scheduleId)
	return nil
}
