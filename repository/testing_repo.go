package repository

import (
	"context"
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/testing"
)

type TestingRepo interface {
	SelectJobByName(context context.Context, jobName string) (model.JobsTesting, error)
	SelectAllUserByJobId(context context.Context, jobId string) ([]model.JobsUser, error)
	SelectGithubByJobId(context context.Context, jobId string) (model.JobsGithub, error)
	SaveActionOfUser(context context.Context, userAction testing.UserActionTesting) (testing.UserActionTesting, error)
	SaveJobTest(context context.Context, job model.JobsTesting) (model.JobsTesting, error)
	RemoveJobTest(context context.Context, jobID string) error
	SaveJobGithub(context context.Context, github model.JobsGithub) (model.JobsGithub, error)
	SaveJobUser(context context.Context, user model.JobsUser) (model.JobsUser, error)
	UpdateAlertTelegramForJob(context context.Context, job model.JobsTesting) (model.JobsTesting, error)
	DisableAlertTelegramJob(context context.Context, AlertTele string) ([]model.JobsTesting, error)
}
