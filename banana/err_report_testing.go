package banana

import "errors"

var (
	//testing
	JobsTestingNotFound = errors.New("Không tìm thấy Jobs Testing")
	RunJobsNotFound     = errors.New("Không tìm thấy Run Jobs")
	RunTestNotFound     = errors.New("Không tìm thấy Run Test")
)
