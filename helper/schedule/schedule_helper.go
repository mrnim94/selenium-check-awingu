package schedule

import (
	"selenium-check-awingu/helper/restclient"
	"selenium-check-awingu/model"
)

type TestingSchedule interface {
	AddJobAtSpecificTime(scheduleT model.ScheduleTesting, restClient restclient.RestClient) error
	RemoveScheduleByTag(scheduleId string) error
}
