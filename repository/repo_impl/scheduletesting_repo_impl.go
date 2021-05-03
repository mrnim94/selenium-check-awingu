package repo_impl

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"selenium-check-awingu/banana"
	"selenium-check-awingu/log"
	"selenium-check-awingu/db"
	"selenium-check-awingu/model"
	"selenium-check-awingu/repository"
	"time"
)

type ScheduleTestingRepoImpl struct {
	sql *db.Sql
}

func NewScheduleTestingRepo(sql *db.Sql) repository.ScheduleTestingRepo {
	return &ScheduleTestingRepoImpl{
		sql: sql,
	}
}

func (st *ScheduleTestingRepoImpl) SaveSignalSchedule(context context.Context, signal model.SignalKey) (model.SignalKey, error)  {
	statement := `
		INSERT INTO signal_schedule(signal_key, created_at, updated_at)
		VALUES(:signal_key, :created_at, :updated_at)
	`
	signal.CreatedAt = time.Now()
	signal.UpdatedAt = time.Now()
	_, err := st.sql.Db.NamedExecContext(context, statement, signal)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return signal, banana.SignalTestingExisted
			}
		}
		return signal, err
	}
	return signal, nil
}

func (st *ScheduleTestingRepoImpl) SelectSignalScheduleBySignId(context context.Context, signalId string) (string, error)  {
	var signalKey string

	err := st.sql.Db.GetContext(context, &signalKey,
		"SELECT signal_key FROM signal_schedule WHERE signal_id = $1", signalId)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err.Error())
			return signalKey, banana.SignalKeyNotFound
		}
		log.Error(err.Error())
		return signalKey, err
	}
	return signalKey, nil
}

func (st *ScheduleTestingRepoImpl) UpdateSignalKey(context context.Context, signalKey model.SignalKey) (string, error)  {
	sqlStatement := `
		UPDATE signal_schedule
		SET 
			signal_key  = :signal_key,
			updated_at = :updated_at
		WHERE signal_id = :signal_id
	`
	signalKey.UpdatedAt = time.Now()
	result, err := st.sql.Db.NamedExecContext(context, sqlStatement, signalKey)
	if err != nil {
		log.Error(err.Error())
		return signalKey.SignalKey, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return signalKey.SignalKey, banana.SignalKeyNotUpdated
	}
	if count == 0 {
		return signalKey.SignalKey, banana.SignalKeyNotUpdated
	}

	return signalKey.SignalKey, nil
}

func (st *ScheduleTestingRepoImpl) SaveScheduleTesting(context context.Context, scheduleT model.ScheduleTesting) (model.ScheduleTesting, error) {
	statement := `
		INSERT INTO selenium_schedule(every, day, at_time, signal_key, job_name, created_at, updated_at)
		VALUES(:every, :day, :at_time, :signal_key, :job_name, :created_at, :updated_at)
	`
	scheduleT.SignalKey = "not_have"
	scheduleT.CreatedAt = time.Now()
	scheduleT.UpdatedAt = time.Now()
	_, err := st.sql.Db.NamedExecContext(context, statement, scheduleT)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return scheduleT, banana.SheduleTestingExisted
			}
		}
		return scheduleT, err
	}
	return scheduleT, nil
}

func (st *ScheduleTestingRepoImpl) SelectAllScheduleDifferenceSignalKey(context context.Context, signalKey string) ([]model.ScheduleTesting, error) {
	scheduleT := []model.ScheduleTesting{}
	err := st.sql.Db.SelectContext(context, &scheduleT,
		`SELECT schedule_id,  every, day, at_time, signal_key, job_name
				FROM selenium_schedule
				WHERE signal_key NOT IN ($1)
		`, signalKey)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err.Error())
			return scheduleT, banana.SignalKeyNotFound
		}
		log.Error(err.Error())
		return scheduleT, err
	}
	return scheduleT, nil

}

func (st *ScheduleTestingRepoImpl) UpdateSignalKeyForSchedule(context context.Context, newSignalKey string, scheduleT model.ScheduleTesting) (model.ScheduleTesting, error) {
	sqlStatement := `
		UPDATE selenium_schedule
		SET 
			signal_key  = :signal_key,
			updated_at = :updated_at
		WHERE schedule_id = :schedule_id
	`
	scheduleT.SignalKey = newSignalKey
	scheduleT.UpdatedAt = time.Now()
	result, err := st.sql.Db.NamedExecContext(context, sqlStatement, scheduleT)
	if err != nil {
		log.Error(err.Error())
		return scheduleT, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return scheduleT, banana.SignalKeyNotUpdated
	}
	if count == 0 {
		return scheduleT, banana.SignalKeyNotUpdated
	}

	return scheduleT, nil
}

func (st *ScheduleTestingRepoImpl)SelectAllScheduleTesting(context context.Context) ([]model.ScheduleTesting, error) {
	scheduleT := []model.ScheduleTesting{}
	err := st.sql.Db.SelectContext(context, &scheduleT,
		`SELECT schedule_id,  every, day, at_time, signal_key, job_name
				FROM selenium_schedule
		`)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err.Error())
			return scheduleT, banana.SignalKeyNotFound
		}
		log.Error(err.Error())
		return scheduleT, err
	}
	return scheduleT, nil
}

func (st *ScheduleTestingRepoImpl) RemoveScheduleById(context context.Context, scheduleId string) error  {
	resultUser := st.sql.Db.MustExecContext(
		context,
		"DELETE FROM selenium_schedule WHERE schedule_id = $1",
		scheduleId)

	_, err := resultUser.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return banana.DeleteScheduleError
	}
	return nil
}