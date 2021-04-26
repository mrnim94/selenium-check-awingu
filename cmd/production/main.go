package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"os"
	"runtime"
	"selenium-check-awingu/db"
	_ "selenium-check-awingu/docs"

	"selenium-check-awingu/handler"
	"selenium-check-awingu/helper"
	"selenium-check-awingu/helper/automate/automate_impl"
	"selenium-check-awingu/log"
	"selenium-check-awingu/repository/repo_impl"
	"selenium-check-awingu/router"
)

func init() {
	os.Setenv("APP_NAME", "backend_testing_pro")
	log.InitLogger(false)
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
	var configPath string
	if runtime.GOOS == "windows" {
		configPath = "./config_file/.env"
	} else {
		configPath = "../../config_file/.env"
	}
	// loads values from .env into the system
	if err := godotenv.Load(configPath); err != nil {
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

	sql.Connect()
	defer sql.Close()

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
	}

	reportTestingHandler := handler.ReportTestingHandler{
		ReportTestingRepo: repo_impl.NewReportTestingRepo(sql),
	}

	api := router.API{
		Echo:                 e,
		AutoTestingHandler:   autoTestingHandler,
		TestingHandler:       testingHandler,
		ReportTestingHandler: reportTestingHandler,
	}
	api.SetupRouter()

	//autoTestingHandler.HandlerTestingConcurrency()

	e.Logger.Fatal(e.Start(":7000"))
}
