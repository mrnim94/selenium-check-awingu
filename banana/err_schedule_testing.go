package banana

import "errors"

var (
	SignalTestingExisted = errors.New("Đã tồn tại Signal Key này")
	SignalKeyNotFound     = errors.New("Không tìm thấy Signal Key")
	SignalKeyNotUpdated = errors.New("Không thể update signal Key")

	SheduleTestingExisted = errors.New("Đã tồn tại Shedule này")

	//restclient
	RunTestingScheduleError = errors.New("Chạy test theo schedule bị lỗi")

	DeleteScheduleError = errors.New("Xóa schedule trong DB bị lỗi")
)
