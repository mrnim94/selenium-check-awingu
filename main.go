package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"os"
	"selenium-check-awingu/db"
	_ "selenium-check-awingu/docs"
	"selenium-check-awingu/handler"
	"selenium-check-awingu/helper"
	"selenium-check-awingu/helper/automate/automate_impl"
	"selenium-check-awingu/helper/restclient/restclient_impl"
	"selenium-check-awingu/helper/schedule"
	"selenium-check-awingu/helper/schedule/schedule_impl"
	"selenium-check-awingu/log"
	"selenium-check-awingu/repository/repo_impl"
	"selenium-check-awingu/router"
	"time"
)

func init() {
	os.Setenv("APP_NAME", "backend_testing_dev")
	log.InitLogger(false)
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
	// loads values from .env into the system
	if err := godotenv.Load("config_file/.env"); err != nil {
		log.Print("No .env file found")
	}
}

// @title Tool Testing API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7000
// @BasePath /
func main() {
	connectStrPostgres, exists := os.LookupEnv("CONNECT_POSTGRES")
	if !(exists) {
		log.Error(exists)
		return
	}
	sql := &db.Sql{
		ConnectString: connectStrPostgres,
	}

	connectStrSelenium, exists := os.LookupEnv("CONNECT_SELENIUM")
	if !(exists) {
		log.Error(exists)
		return
	}
	seConNim := &automate_impl.Selenium{
		Browser:       "chrome",
		ConnectServer: connectStrSelenium,
	}

	scheduleNim := &schedule.GoCron{
		ZoneName:     "VietNam",
		SecondsOfUTC: 7*60*60,
	}

	connectMyApi, exists := os.LookupEnv("CONNECT_MY_API")
	if !(exists) {
		log.Error(exists)
		return
	}
	urlNim := restclient_impl.Resty{Url: connectMyApi}

	//Database
	sql.Connect()
	defer sql.Close()

	//Scheduler
	scheduleNim.GetGoCron()

	urlTelegramApi := restclient_impl.Resty{Url: "https://api.telegram.org"}

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()
	e.Validator = structValidator

	autoTestingHandler := handler.AutoTestingHandler{
		UserTesting: repo_impl.NewUserTesting(sql),
		Automate:    automate_impl.NewSelenium(seConNim),
	}

	testingHandler := handler.TestingHandler{
		TestingRepo: repo_impl.NewTestingRepo(sql),
		Automate:    automate_impl.NewSelenium(seConNim),
		AlertRepo: repo_impl.NewAlertRepo(sql),
		RestClient: restclient_impl.NewResty(urlTelegramApi),
	}

	reportTestingHandler := handler.ReportTestingHandler{
		ReportTestingRepo: repo_impl.NewReportTestingRepo(sql),
	}

	scheduleTestingHandler := handler.ScheduleTestingHandler{
		ScheduleTestingRepo : repo_impl.NewScheduleTestingRepo(sql),
		ScheduleHelper: schedule_impl.NewTestingSchedule(scheduleNim),
		RestClientHelper: restclient_impl.NewResty(urlNim),
	}

	alertHandler := handler.AlertHandler{
		AlertRepo: repo_impl.NewAlertRepo(sql),
		TestingRepo: repo_impl.NewTestingRepo(sql),
	}

	api := router.API{
		Echo:                 e,
		AutoTestingHandler:   autoTestingHandler,
		TestingHandler:       testingHandler,
		ReportTestingHandler: reportTestingHandler,
		ScheduleTestingHandler: scheduleTestingHandler,
		AlertHandler: alertHandler,
	}
	api.SetupRouter()

	scheduleTestingHandler.HandlerSignalSchedule()
	//scheduleTestingHandler.HandlerScheduleAtSpecificTime()

	go intervalScanSchedule(30*time.Second, scheduleTestingHandler)

	e.Logger.Fatal(e.Start(":7000"))
}

func intervalScanSchedule(timeSchedule time.Duration, handler handler.ScheduleTestingHandler) {
	ticker := time.NewTicker(timeSchedule)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Scan schedule...")
				handler.HandlerScheduleAtSpecificTime()
			}
		}
	}()
}
