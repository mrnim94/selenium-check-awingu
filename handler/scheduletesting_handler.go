package handler

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"selenium-check-awingu/banana"
	"selenium-check-awingu/helper/restclient"
	"selenium-check-awingu/helper/schedule"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/req"
	"selenium-check-awingu/repository"
	"time"
)

type ScheduleTestingHandler struct {
	ScheduleTestingRepo repository.ScheduleTestingRepo
	ScheduleHelper schedule.TestingSchedule
	RestClientHelper restclient.RestClient
}

func (st *ScheduleTestingHandler) HandlerSignalSchedule() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	signalKey, err := st.ScheduleTestingRepo.SelectSignalScheduleBySignId(ctx, "1")
	if err != nil {
		log.Error(err.Error())
		signKey, err := uuid.NewUUID()
		if err != nil {
			log.Error(err.Error())
			if err == banana.UserBookingNotFound {
				log.Error(err.Error())
				return err
			}
		}
		signalKey := model.SignalKey{
			SignalKey: signKey.String(),
		}
		signalKey, err = st.ScheduleTestingRepo.SaveSignalSchedule(ctx, signalKey)
		if err != nil {
			log.Error(err.Error())
			if err == banana.UserBookingNotFound {
				log.Error(err.Error())
				return err
			}
		}
		log.Info("Signal Key lần đâu tiên được tạo là: "+ signalKey.SignalKey)
	}else {
		log.Info("Signal Key "+signalKey+" sẽ được thay đổi")
		signKey, err := uuid.NewUUID()
		if err != nil {
			log.Error(err.Error())
			if err == banana.UserBookingNotFound {
				log.Error(err.Error())
				return err
			}
		}
		signalkeyUpdate := model.SignalKey{
			SignalKey: signKey.String(),
			SignalId: "1",
		}
		newSignalKey, err := st.ScheduleTestingRepo.UpdateSignalKey(ctx, signalkeyUpdate)
		if err != nil {
			log.Error(err.Error())
			if err == banana.UserBookingNotFound {
				log.Error(err.Error())
				return err
			}
		}
		log.Info("Signal Key đã được cập nhật là: "+ newSignalKey)
	}
	return nil
}


// RunTesting godoc
// @Summary Add schedule để check testing
// @Tags schedule-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestAddSchedule true "schedule"
// @Success 200 {object} model.Response
// @Failure 409 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /schedule/add [post]
func (st *ScheduleTestingHandler) HandlerAddSchedule(c echo.Context) error {
	req := req.RequestAddSchedule{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("Lấy được request Add Schedule")
	}

	if err := c.Validate(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("kiểm trả các thông số gửi lên cho Add Schedule thành công")
	}

	newSchedule := model.ScheduleTesting{
		Every:     req.Every,
		Day:       req.Day,
		AtTime:    req.AtTime,
		JobName:   req.JobName,
	}

	schedule, err := st.ScheduleTestingRepo.SaveScheduleTesting(c.Request().Context(), newSchedule)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       schedule,
	})
}

func (st *ScheduleTestingHandler) HandlerScheduleAtSpecificTime() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	signalKey, err := st.ScheduleTestingRepo.SelectSignalScheduleBySignId(ctx, "1")
	if err != nil {
		log.Error(err.Error())
		return err
	}

	listSchedules, err := st.ScheduleTestingRepo.SelectAllScheduleDifferenceSignalKey(ctx, signalKey)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	cSchedule := make(chan model.ScheduleTesting)
	done := make(chan bool)
	go func() {
		for _, schedule := range listSchedules {
			cSchedule <- schedule
		}
		done <- true
	}()

	for i:=0; i <= len(listSchedules); i++ {
		select {
		case resultSchedule := <- cSchedule:
			log.Info("Schedule " + resultSchedule.ScheduleId + " sẽ được cập nhật Signal Key " + signalKey)
			scheduleN ,err := st.ScheduleTestingRepo.UpdateSignalKeyForSchedule(ctx, signalKey, resultSchedule)
			if err != nil {
				log.Error(err.Error())
			}else {
				log.Info("Schedule " + resultSchedule.ScheduleId + " sẽ được ghi vào Schedule")
				st.ScheduleHelper.AddJobAtSpecificTime(scheduleN, st.RestClientHelper)
				if err != nil {
					log.Error(err.Error())
				}
			}
		case <-done:
			//log.Info("Done Scan schedule")
			fmt.Println("Done Scan schedule")
		}
	}
	return nil
}

// RunTesting godoc
// @Summary Hiện thị các schedule
// @Tags schedule-service
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Response
// @Failure 409 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /schedule/list [get]
func (st *ScheduleTestingHandler) HandlerListSchedules(c echo.Context) error  {
	schedules, err := st.ScheduleTestingRepo.SelectAllScheduleTesting(c.Request().Context())
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       schedules,
	})
}

// RunTesting godoc
// @Summary Xóa schedule để check testing
// @Tags schedule-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestDeleteSchedule true "schedule"
// @Success 200 {object} model.Response
// @Failure 409 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /schedule/delete [post]
func (st *ScheduleTestingHandler) HandlerRemoveSchedules(c echo.Context) error  {
	reqDS := req.RequestDeleteSchedule{}

	if err := c.Bind(&reqDS); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("Lấy được request Id Schedule")
	}

	if err := c.Validate(reqDS); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err := st.ScheduleHelper.RemoveScheduleByTag(reqDS.ScheduleId)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err = st.ScheduleTestingRepo.RemoveScheduleById(c.Request().Context(), reqDS.ScheduleId)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       reqDS,
	})

}

