package banana

import "errors"

var (
	//testing
	JobTestingNotFound = errors.New("Không tìm thấy Job Testing")
	JobTestingExisted = errors.New("Đã tồn tại Job này")
	DelJobFail = errors.New("Xóa Job bị lỗi")

	JobUserNotFound    = errors.New("Không tìm thấy Job User")
	JobUserExisted = errors.New("Đã tồn tại User này")
	DelUserFail = errors.New("Xóa User bị lỗi")

	JobGithubNotFound  = errors.New("Không tìm thấy Thông tin Github")
	JobGithubExisted = errors.New("Đã tồn tại Github này")
	DelGithubFail = errors.New("Xóa Github bị lỗi")

	ActionConflict     = errors.New("Action bị trùng")
	ActionFail         = errors.New("Ghi nhận Action thất bại")
)
