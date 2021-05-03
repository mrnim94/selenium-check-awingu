package router

import (
	"github.com/labstack/echo/v4"
	"selenium-check-awingu/handler"
)

type API struct {
	Echo                 *echo.Echo
	AutoTestingHandler   handler.AutoTestingHandler
	TestingHandler       handler.TestingHandler
	ReportTestingHandler handler.ReportTestingHandler
	ScheduleTestingHandler handler.ScheduleTestingHandler
	AlertHandler handler.AlertHandler
}

func (api *API) SetupRouter() {
	api.Echo.GET("/", handler.Welcome)

	// api Sử dụng cho Testing
	api.Echo.POST("/tester/run-testing", api.TestingHandler.HandlerRunTesting)
	api.Echo.POST("/tester/add-job", api.TestingHandler.HandlerAddJob)
	api.Echo.DELETE("/tester/delete-job/:jobid", api.TestingHandler.HandlerDeleteJob)
	api.Echo.POST("/tester/add-github", api.TestingHandler.HandlerAddGithubJob)
	api.Echo.POST("/tester/add-user", api.TestingHandler.HandlerAddUserJob)
	api.Echo.POST("/tester/update-alert-telegram", api.TestingHandler.HandlerUpdateAlertTelegramForJob)



	// api Sử dụng cho Report Testing
	api.Echo.GET("/report/jobs-testing", api.ReportTestingHandler.HandlerSelectJobsTesting)
	api.Echo.POST("/report/run-jobs", api.ReportTestingHandler.HandlerSelectRunJobs)
	api.Echo.POST("/report/run-test", api.ReportTestingHandler.HandlerSelectRunTest)

	//schedule
	schedule := api.Echo.Group("/schedule")
	schedule.POST("/add", api.ScheduleTestingHandler.HandlerAddSchedule)
	schedule.GET("/list", api.ScheduleTestingHandler.HandlerListSchedules)
	schedule.POST("/delete", api.ScheduleTestingHandler.HandlerRemoveSchedules)

	//alert
	alert := api.Echo.Group("/alert")
	alert.POST("/add-telegram", api.AlertHandler.HandlerAddTelegramInfo)
	alert.GET("/list-telegram", api.AlertHandler.HandlerListTelegramInfo)
	alert.POST("/delete-telegram", api.AlertHandler.HandlerDeleteTelegramInfo)

}
