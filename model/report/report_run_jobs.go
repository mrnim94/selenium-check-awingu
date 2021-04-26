package report

type ReportRunJobs struct {
	JobId     string `json:"jobId" db:"job_id, omitempty"`
	TestId    string `json:"testId" db:"test_id, omitempty"`
	NameTest  string `json:"nameTest" db:"name_test, omitempty"`
	Version   string `json:"version" db:"version, omitempty"`
	Browser   string `json:"browser" db:"browser, omitempty"`
	Agent     string `json:"agent" db:"agent, omitempty"`
	Status    string `json:"status" db:"status, omitempty"`
	StartTime int64  `json:"startTime" db:"start_time, omitempty"`
	EndTime   int64  `json:"endTime" db:"end_time, omitempty"`
}
