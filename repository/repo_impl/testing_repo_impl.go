package repo_impl

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"selenium-check-awingu/banana"
	"selenium-check-awingu/db"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/testing"
	"selenium-check-awingu/repository"
	"time"
)

type TestingRepoImpl struct {
	sql *db.Sql
}

func NewTestingRepo(sql *db.Sql) repository.TestingRepo {
	return &TestingRepoImpl{
		sql: sql,
	}
}

func (t *TestingRepoImpl) SelectJobByName(context context.Context, jobName string) (model.JobsTesting, error) {
	var job model.JobsTesting
	err := t.sql.Db.GetContext(context, &job,
		"SELECT * FROM jobs_testing WHERE job_name = $1 AND status = '1'", jobName)

	if err != nil {
		if err == sql.ErrNoRows {
			return job, banana.JobTestingNotFound
		}
		log.Error(err.Error())
		return job, err
	}

	return job, nil
}

func (t *TestingRepoImpl) SelectAllUserByJobId(context context.Context, jobId string) ([]model.JobsUser, error) {
	var users []model.JobsUser
	err := t.sql.Db.SelectContext(context, &users,
		"SELECT * FROM jobs_user WHERE job_id = $1 AND status = '1'", jobId)

	if err != nil {
		if err == sql.ErrNoRows {
			return users, banana.JobUserNotFound
		}
		log.Error(err.Error())
		return users, err
	}

	return users, nil
}

func (t *TestingRepoImpl) SelectGithubByJobId(context context.Context, jobId string) (model.JobsGithub, error) {
	var github model.JobsGithub
	err := t.sql.Db.GetContext(context, &github,
		"SELECT * FROM jobs_github WHERE job_id = $1", jobId)

	if err != nil {
		if err == sql.ErrNoRows {
			return github, banana.JobGithubNotFound
		}
		log.Error(err.Error())
		return github, err
	}

	return github, nil
}

func (u *TestingRepoImpl) SaveActionOfUser(context context.Context, userAction testing.UserActionTesting) (testing.UserActionTesting, error) {
	statement := `
		INSERT INTO jobs_log(job_id, time_id, test_id, name_test, version, browser, page, agent, web_driver, description, action, data, data1)
		VALUES(:job_id, :time_id, :test_id, :name_test, :version, :browser, :page, :agent, :web_driver, :description, :action, :data, :data1)
	`
	_, err := u.sql.Db.NamedExecContext(context, statement, userAction)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return userAction, banana.ActionConflict
			}
		}
		return userAction, banana.ActionFail
	}

	return userAction, nil
}

func (u *TestingRepoImpl) SaveJobTest(context context.Context, job model.JobsTesting) (model.JobsTesting, error) {
	statement := `
		INSERT INTO jobs_testing(job_id, job_name, status, created_at, updated_at)
		VALUES(:job_id, :job_name, :status, :created_at, :updated_at)
	`

	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()
	_, err := u.sql.Db.NamedExecContext(context, statement, job)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return job, banana.JobTestingExisted
			}
		}
		return job, err
	}

	return job, nil
}

func (u *TestingRepoImpl) SaveJobGithub(context context.Context, github model.JobsGithub) (model.JobsGithub, error) {
	statement := `
		INSERT INTO jobs_github(access_token, owner, repo, path, job_id, created_at, updated_at)
		VALUES(:access_token, :owner, :repo, :path, :job_id, :created_at, :updated_at)
	`

	github.CreatedAt = time.Now()
	github.UpdatedAt = time.Now()
	_, err := u.sql.Db.NamedExecContext(context, statement, github)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return github, banana.JobGithubExisted
			}
		}
		return github, err
	}

	return github, nil
}

func (u *TestingRepoImpl) SaveJobUser(context context.Context, user model.JobsUser) (model.JobsUser, error) {
	statement := `
		INSERT INTO jobs_user(username, password, job_id, status, created_at, updated_at)
		VALUES(:username, :password, :job_id, :status, :created_at, :updated_at)
	`

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, banana.JobUserExisted
			}
		}
		return user, err
	}

	return user, nil
}

func (u *TestingRepoImpl) RemoveJobTest(context context.Context, jobID string) error {
	resultUser := u.sql.Db.MustExecContext(
		context,
		"DELETE FROM jobs_user WHERE job_id = $1",
		jobID)

	_, err := resultUser.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return banana.DelUserFail
	}

	resultGithub := u.sql.Db.MustExecContext(
		context,
		"DELETE FROM jobs_github WHERE job_id = $1",
		jobID)

	_, err = resultGithub.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return banana.DelGithubFail
	}

	resultJob := u.sql.Db.MustExecContext(
		context,
		"DELETE FROM jobs_testing WHERE job_id = $1",
		jobID)

	_, err = resultJob.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return banana.DelJobFail
	}

	return nil
}
