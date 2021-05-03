package repository

import (
	"selenium-check-awingu/model/alert"
	"context"
)

type AlertRepo interface {
	SaveTelegramInfo(context context.Context, tele alert.TelegramInfo) (alert.TelegramInfo, error)
	SelectAllTelegram(context context.Context) ([]alert.TelegramInfo, error)
	SelectTelegramInfoByName(context context.Context, telegramName string) (alert.TelegramInfo, error)
	RemoveTelegramInfoByName(context context.Context, telegramName string) error
}
