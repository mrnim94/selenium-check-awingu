package repo_impl

import (
	"context"
	"database/sql"
	"selenium-check-awingu/banana"
	"selenium-check-awingu/db"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model/report"
	"selenium-check-awingu/model/testing"
	"selenium-check-awingu/repository"
)

type ReportTestingRepoImpl struct {
	sql *db.Sql
}

func NewReportTestingRepo(sql *db.Sql) repository.ReportTestingRepo {
	return &ReportTestingRepoImpl{
		sql: sql,
	}
}

func (rt *ReportTestingRepoImpl) SelectAllJobsTesting(context context.Context, jobId string) ([]report.ReportJobsTesting, error) {
	jobsT := []report.ReportJobsTesting{}
	//log.Info("thực hiện lấy " + jobId + " Jobs Testing trong DB")
	err := rt.sql.Db.SelectContext(context, &jobsT,
		`SELECT jobs_testing.job_id, jobs_testing.job_name, jobs_testing.status, jobs_testing.alert_telegram,
						jobs_github.owner, jobs_github.repo, jobs_github.path
				FROM jobs_testing
				INNER JOIN jobs_github
				ON jobs_testing.job_id = jobs_github.job_id
		`) // true as jobs_github nếu bẳng jobs_github thì nó trả lại là true

	if err != nil {
		if err == sql.ErrNoRows {
			return jobsT, banana.JobsTestingNotFound
		}
		log.Error(err.Error())
		return jobsT, err
	}
	return jobsT, nil
}

func (rt *ReportTestingRepoImpl) SelectRunJobsByJobId(context context.Context,
	jobId string, startTime, endTime int64) ([]report.ReportRunJobs, error) {
	runJobs := []report.ReportRunJobs{}
	err := rt.sql.Db.SelectContext(context, &runJobs,
		`SELECT job_id, test_id, name_test, version, browser, agent, time_id AS start_time, ('9223372036854775800') as end_time 
				FROM jobs_log 
				WHERE page = 'startTesting' 
				AND action = 'Begin'
				AND job_id = $1
				AND time_id >= $2
				AND 
				test_id NOT IN (
					SELECT st.test_id 
					FROM 
						(SELECT time_id AS start_time, test_id, name_test, version, browser, page, agent 
						FROM jobs_log 
						WHERE page = 'startTesting' 
						AND action = 'Begin'
						AND job_id = $1
						AND time_id >= $2) AS st 
					INNER JOIN 
						(SELECT time_id AS end_time, test_id 
						FROM jobs_log 
						WHERE page = 'endTesting' 
						AND action = 'End'
						AND job_id = $1
						AND time_id <= $3) AS et 
					ON st.test_id = et.test_id)
				UNION
				SELECT * FROM(SELECT st.job_id, st.test_id, st.name_test, st.version, st.browser, st.agent, st.start_time, et.end_time 
								FROM (SELECT time_id AS start_time, job_id, test_id, name_test, version, browser, page, agent 
										FROM jobs_log 
										WHERE page = 'startTesting' 
										AND action = 'Begin'
										AND job_id = $1
										AND time_id >= $2) AS st 
								INNER JOIN (SELECT time_id 
											AS end_time, test_id 
											FROM jobs_log 
											WHERE page = 'endTesting' 
											AND action = 'End'
											AND job_id = $1
											AND time_id <= $3) AS et 
								ON st.test_id = et.test_id) AS nim
				ORDER BY start_time DESC
		`, jobId, startTime, endTime)

	if err != nil {
		if err == sql.ErrNoRows {
			return runJobs, banana.RunJobsNotFound
		}
		log.Error(err.Error())
		return runJobs, err
	}
	return runJobs, nil
}

func (rt *ReportTestingRepoImpl) SelectRunTestByTestId(context context.Context, testId string) ([]testing.UserActionTesting, error) {
	runTest := []testing.UserActionTesting{}
	err := rt.sql.Db.SelectContext(context, &runTest, `
			SELECT * FROM jobs_log WHERE test_id = $1 ORDER BY time_id ASC
	`, testId)

	if err != nil {
		if err == sql.ErrNoRows {
			return runTest, banana.RunTestNotFound
		}
		log.Error(err.Error())
		return runTest, err
	}
	return runTest, nil
}

func (rt *ReportTestingRepoImpl) CheckStatusByTestId(context context.Context, testId string) (string, error) {
	var statusTestId string

	err := rt.sql.Db.GetContext(context, &statusTestId,
		"SELECT action FROM jobs_log WHERE test_id = $1 AND action = 'Error'", testId)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err.Error() + " status Error")
			err := rt.sql.Db.GetContext(context, &statusTestId,
				"SELECT action FROM jobs_log WHERE test_id = $1 AND action = 'Warning'", testId)
			if err != nil {
				if err == sql.ErrNoRows {
					log.Error(err.Error() + " status Warning")
					statusTestId = "Not_Issue"
					return statusTestId, nil
				}
				log.Error(err.Error())
				return statusTestId, err
			}

			return statusTestId, nil
		}
		log.Error(err.Error())
		return statusTestId, err
	}

	return statusTestId, nil
}
