package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/alert"
	"selenium-check-awingu/model/req"
	"selenium-check-awingu/repository"
)

type AlertHandler struct {
	AlertRepo repository.AlertRepo
	TestingRepo repository.TestingRepo
}

// RunTesting godoc
// @Summary Add thông tin telegram để alert
// @Tags Alert-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestInfoTelegram true "alert"
// @Success 200 {object} model.Response
// @Failure 409 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /alert/add-telegram [post]
func (al *AlertHandler) HandlerAddTelegramInfo(c echo.Context) error {
	req := req.RequestInfoTelegram{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("Lấy được request Add Telegram Info")
	}

	if err := c.Validate(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("kiểm trả các thông số gửi lên cho Add Telegram Info thành công")
	}

	newTelegramInfo := alert.TelegramInfo{
		TelegramName: req.TelegramName,
		TelegramToken:       req.TelegramToken,
		ChatId:              req.ChatId,
		DisableNotification: req.DisableNotification,
	}

	tele, err :=  al.AlertRepo.SaveTelegramInfo(c.Request().Context(), newTelegramInfo)
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
		Data:       tele,
	})
}

// RunTesting godoc
// @Summary Hiển thị thông tin telegram để alert
// @Tags Alert-service
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Response
// @Failure 409 {object} model.Response
// @Router /alert/list-telegram [get]
func (al *AlertHandler) HandlerListTelegramInfo(c echo.Context) error {

	listTelegram, err := al.AlertRepo.SelectAllTelegram(c.Request().Context())
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
		Data:       listTelegram,
	})
}

// RunTesting godoc
// @Summary xóa thông tin telegram để alert
// @Tags Alert-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestDeleteTelegram true "alert"
// @Success 200 {object} model.Response
// @Failure 409 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /alert/delete-telegram [post]
func (al *AlertHandler) HandlerDeleteTelegramInfo(c echo.Context) error {
	req := req.RequestDeleteTelegram{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("Lấy được request Delete Telegram Info")
	}

	if err := c.Validate(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("kiểm trả các thông số gửi lên cho Delete Telegram Info thành công")
	}

	jobs, err := al.TestingRepo.DisableAlertTelegramJob(c.Request().Context(), req.TelegramName)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	fmt.Println(jobs)

	err = al.AlertRepo.RemoveTelegramInfoByName(c.Request().Context(), req.TelegramName)
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
		Data:       jobs,
	})
}