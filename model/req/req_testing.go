package req

type RequestTesting struct {
	JobName string `json:"jodName,omitempty" validate:"required"`
}

type RequestRunJobs struct {
	JobId     string `json:"jobId,omitempty" validate:"required"`
	StartTime int64  `json:"startTime,omitempty" validate:"required"`
	EndTime   int64  `json:"endTime,omitempty" validate:"required"`
}

type RequestRunTest struct {
	TestId string `json:"testId,omitempty" validate:"required"`
}

type RequestAddJob struct {
	JobName   string    `json:"jobName,omitempty" validate:"required"`
	Status    int    `json:"status,omitempty" validate:"required"`
}

type RequestDeleteJob struct {
	JobName   string    `json:"jobName,omitempty" validate:"required"`
}

type RequestAddGithubJob struct {
	AccessToken string    `json:"accessToken,omitempty" validate:"required"`
	Owner       string    `json:"owner,omitempty" validate:"required"`
	Repo        string    `json:"repo,omitempty" validate:"required"`
	Path        string    `json:"path,omitempty" validate:"required"`
	JobID       string    `json:"jobID,omitempty" validate:"required"`
}

type RequestAddUserJob struct {
	Username  string    `json:"userName,omitempty" validate:"required"`
	Password  string    `json:"password,omitempty" validate:"required"`
	JobId     string    `json:"jobId,omitempty" validate:"required"`
	Status    int    `json:"status,omitempty" validate:"required"`
}

type RequestUpdateAlertTelegramJob struct {
	JobId     string `json:"jobId,omitempty" validate:"required"`
	AlertTelegram   string    `json:"alertTelegram,omitempty" validate:"required"`
}
