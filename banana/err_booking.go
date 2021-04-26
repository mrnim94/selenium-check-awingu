package banana

import "errors"

var (
	//booking
	UserBookingNotFound = errors.New("Không tìm thấy User")
	UserNotUpdated      = errors.New("Không update được User")
)
