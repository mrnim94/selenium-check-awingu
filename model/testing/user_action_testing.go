package testing

type UserActionTesting struct {
	Id          string `json:"-" db:"id, omitempty"`
	JobId       string `json:"jobId" db:"job_id, omitempty"`
	TimeId      int64  `json:"timeId" db:"time_id, omitempty"`
	TestId      string `json:"testId" db:"test_id, omitempty"`
	NameTest    string `json:"nameTest" db:"name_test, omitempty"`
	Version     string `json:"version" db:"version, omitempty"`
	Browser     string `json:"browser" db:"browser, omitempty"`
	Page        string `json:"page" db:"page, omitempty"`
	Agent       string `json:"agent" db:"agent, omitempty"`
	WebDriver   string `json:"webDriver" db:"web_driver, omitempty"`
	Description string `json:"description" db:"description, omitempty"`
	Action      string `json:"action" db:"action, omitempty"`
	Data        string `json:"data" db:"data, omitempty"`
	Data1       string `json:"data1" db:"data1, omitempty"`
}
