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
}

func (api *API) SetupRouter() {
	api.Echo.GET("/", handler.Welcome)

	// api Sử dụng cho Testing
	api.Echo.POST("/tester/run-testing", api.TestingHandler.HandlerRunTesting)
	api.Echo.POST("/tester/add-job", api.TestingHandler.HandlerAddJob)
	api.Echo.DELETE("/tester/delete-job/:jobid", api.TestingHandler.HandlerDeleteJob)

	api.Echo.POST("/tester/add-github", api.TestingHandler.HandlerAddGithubJob)

	api.Echo.POST("/tester/add-user", api.TestingHandler.HandlerAddUserJob)



	// api Sử dụng cho Report Testing
	api.Echo.GET("/report/jobs-testing", api.ReportTestingHandler.HandlerSelectJobsTesting)
	api.Echo.POST("/report/run-jobs", api.ReportTestingHandler.HandlerSelectRunJobs)
	api.Echo.POST("/report/run-test", api.ReportTestingHandler.HandlerSelectRunTest)
}
