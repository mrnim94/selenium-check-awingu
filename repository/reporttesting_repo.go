package repository

import (
	"context"
	"selenium-check-awingu/model/report"
	"selenium-check-awingu/model/testing"
)

type ReportTestingRepo interface {
	SelectAllJobsTesting(context context.Context, jobId string) ([]report.ReportJobsTesting, error)
	SelectRunJobsByJobId(context context.Context, jobId string, startTime, endTime int64) ([]report.ReportRunJobs, error)
	SelectRunTestByTestId(context context.Context, testId string) ([]testing.UserActionTesting, error)
	CheckStatusByTestId(context context.Context, testId string) (string, error)
}
