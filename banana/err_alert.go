package banana

import "errors"

var (
	TelegramInfoExisted = errors.New("Đã tồn tại Telegram này, chọn tên khác")
	DeleteTelegramInfoError = errors.New("Xóa Telegram trong DB bị lỗi")
)
