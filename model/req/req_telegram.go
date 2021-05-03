package req

type RequestSendMessageToGroup struct {
	ChatId string `json:"chat_id,omitempty" validate:"required"`
	Text string `json:"text,omitempty" validate:"required"`
	DisableNotification bool `json:"disable_notification,omitempty" validate:"required"`
}


type RequestInfoTelegram struct {
	TelegramName string `json:"telegramName,omitempty" validate:"required"`
	TelegramToken string `json:"telegramToken,omitempty" validate:"required"`
	ChatId string `json:"chatId,omitempty" validate:"required"`
	DisableNotification string `json:"disableNotification,omitempty" validate:"required"`
}

type RequestDeleteTelegram struct {
	TelegramName string `json:"telegramName,omitempty" validate:"required"`
}