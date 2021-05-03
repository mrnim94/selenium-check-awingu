package alert

import "time"

type TelegramInfo struct {
	TelegramId string `json:"telegramId,omitempty" db:"telegram_id, omitempty"`
	TelegramName string `json:"telegramName,omitempty" db:"telegram_name, omitempty"`
	TelegramToken string `json:"telegramToken,omitempty" db:"telegram_token, omitempty"`
	ChatId string `json:"chatId,omitempty" db:"chat_id, omitempty"`
	DisableNotification string `json:"disableNotification,omitempty" db:"disable_notification, omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at, omitempty"`
}