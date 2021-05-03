package repo_impl

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"selenium-check-awingu/banana"
	"selenium-check-awingu/db"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model/alert"
	"selenium-check-awingu/repository"
	"time"
)

type AlertRepoImpl struct {
	sql *db.Sql
}

func NewAlertRepo(sql *db.Sql) repository.AlertRepo {
	return &AlertRepoImpl{
		sql: sql,
	}
}

func (st *AlertRepoImpl) SaveTelegramInfo(context context.Context, tele alert.TelegramInfo) (alert.TelegramInfo, error)  {
	statement := `
		INSERT INTO telegram_info(telegram_name, telegram_token, chat_id, disable_notification, created_at, updated_at)
		VALUES(:telegram_name, :telegram_token, :chat_id, :disable_notification, :created_at, :updated_at)
	`
	tele.CreatedAt = time.Now()
	tele.UpdatedAt = time.Now()
	_, err := st.sql.Db.NamedExecContext(context, statement, tele)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return tele, banana.TelegramInfoExisted
			}
		}
		return tele, err
	}
	return tele, nil
}

func (st *AlertRepoImpl) SelectAllTelegram(context context.Context) ([]alert.TelegramInfo, error) {
	listTelegram := []alert.TelegramInfo{}

	err := st.sql.Db.SelectContext(context, &listTelegram,
		`SELECT telegram_id,  telegram_name, telegram_token, chat_id, disable_notification
				FROM telegram_info
		`)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err.Error())
			return listTelegram, banana.SignalKeyNotFound
		}
		log.Error(err.Error())
		return listTelegram, err
	}
	return listTelegram, nil
}

func (st *AlertRepoImpl) SelectTelegramInfoByName(context context.Context, telegramName string) (alert.TelegramInfo, error)  {
	var alert alert.TelegramInfo
	err := st.sql.Db.GetContext(context, &alert,
		"SELECT * FROM telegram_info WHERE telegram_name = $1", telegramName)

	if err != nil {
		if err == sql.ErrNoRows {
			return alert, banana.JobGithubNotFound
		}
		log.Error(err.Error())
		return alert, err
	}

	return alert, nil
}

func (st *AlertRepoImpl) RemoveTelegramInfoByName(context context.Context, telegramName string) error  {
	resultUser := st.sql.Db.MustExecContext(
		context,
		"DELETE FROM telegram_info WHERE telegram_name = $1",
		telegramName)

	_, err := resultUser.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return banana.DeleteTelegramInfoError
	}
	return nil
}

