package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/req"
	"selenium-check-awingu/repository"
)

type ReportTestingHandler struct {
	ReportTestingRepo repository.ReportTestingRepo
}

const Message200ok = "This should be a constant"

// SelectJobsTesting godoc
// @Summary Hiện thị các Job Testing
// @Tags report-testing-service
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Response
// @Success 404 {object} model.Response
// @Router /report/jobs-testing [get]
func (rt *ReportTestingHandler) HandlerSelectJobsTesting(c echo.Context) error {
	jobsResult, err := rt.ReportTestingRepo.SelectAllJobsTesting(c.Request().Context(), "all")

	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    Message200ok,
		Data:       jobsResult,
	})
}

// SelectRunJobs godoc
// @Summary Hiện thị các Testing chạy trong 1 Job
// @Tags report-testing-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestRunJobs true "testing"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Success 409 {object} model.Response
// @Router /report/run-jobs [post]
func (rt *ReportTestingHandler) HandlerSelectRunJobs(c echo.Context) error {
	reqSRJ := req.RequestRunJobs{}
	if err := c.Bind(&reqSRJ); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	if err := c.Validate(reqSRJ); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	runJobs, err := rt.ReportTestingRepo.SelectRunJobsByJobId(c.Request().Context(), reqSRJ.JobId, reqSRJ.StartTime, reqSRJ.EndTime)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	for i, e := range runJobs {
		statusTest, _ := rt.ReportTestingRepo.CheckStatusByTestId(c.Request().Context(), e.TestId)
		if err != nil {
			return c.JSON(http.StatusConflict, model.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
				Data:       nil,
			})
		}
		e.Status = statusTest
		runJobs[i] = e
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    Message200ok,
		Data:       runJobs,
	})
}

// SelectRunTest godoc
// @Summary Hiện thị chi tiết các bước chạy trong 1 testing
// @Tags report-testing-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestRunTest true "testing"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Success 409 {object} model.Response
// @Router /report/run-test [post]
func (rt *ReportTestingHandler) HandlerSelectRunTest(c echo.Context) error {
	reqRRT := req.RequestRunTest{}
	if err := c.Bind(&reqRRT); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	if err := c.Validate(reqRRT); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	runTest, err := rt.ReportTestingRepo.SelectRunTestByTestId(c.Request().Context(), reqRRT.TestId)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    Message200ok,
		Data:       runTest,
	})
}
