package req

type RequestAddSchedule struct {
	Every   string `json:"every,omitempty" validate:"required"`
	Day     string `json:"day,omitempty" validate:"required"`
	AtTime  string `json:"atTime,omitempty" validate:"required"`
	JobName string `json:"jobName,omitempty" validate:"required"`
	//SignalKey string `json:"signalKey,omitempty" validate:"required"`
}

type RequestDeleteSchedule struct {
	ScheduleId string `json:"scheduleId" validate:"required"`
}
