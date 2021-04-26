package automate_impl

import (
	"fmt"
	uuid "github.com/google/uuid"
	"reflect"
	"selenium-check-awingu/helper/automate/automate_impl/automate_plugin"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/testing"
	"selenium-check-awingu/repository"
	"strconv"
)

func (s *Selenium) RobotAutoImpl(contentTesting testing.YamlTesting, userSignIn model.JobsUser, testingRepo repository.TestingRepo, jobID string) error {
	conctRemote, err := s.getRemote()
	if err != nil {
		log.Error(err.Error() + ">>>" + userSignIn.Username)
		return err
	} else {
		log.Info(contentTesting.NameTest + " " + userSignIn.Username + " đã connect thành công đến selenium")
	}
	defer conctRemote.Quit()

	err = conctRemote.Get(contentTesting.Web)
	if err != nil {
		log.Error(err.Error())
		return err
	} else {
		log.Info(contentTesting.NameTest + " " + userSignIn.Username + " đã truy cập được website")
	}

	//width, err := strconv.Atoi(contentTesting.ResizeWindow["width"])
	//if err != nil {
	//	log.Error(err.Error())
	//	return err
	//}
	//height, err := strconv.Atoi(contentTesting.ResizeWindow["height"])
	//if err != nil {
	//	log.Error(err.Error())
	//	return err
	//}
	//err = conctRemote.ResizeWindow("thangta", 1280, 720)
	//if err != nil {
	//	log.Error(err.Error())
	//	return err
	//}

	testId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	for iHook, hook := range contentTesting.Hooks {
		if reflect.ValueOf(hook).Kind() == reflect.Struct {
			vHook := reflect.ValueOf(hook) // ờ đây chắc chắn là v.Kind() là truct
			log.Info(contentTesting.NameTest + " >>> Thực hiện testing trên page " + strconv.Itoa(iHook) + " " + fmt.Sprintf("%v", vHook.Field(0)) + ": " + fmt.Sprintf("%v", vHook.Field(0)))
			steps := reflect.ValueOf(vHook.Field(1).Interface())
			if steps.Kind() == reflect.Slice {
				for t := 0; t < steps.Len(); t++ {
					step := reflect.ValueOf(steps.Index(t).Interface())
					var webDriver, description, webElement, actions, timeout, tab, script, newUrl, checkElements reflect.Value
					if step.Kind() == reflect.Slice {
						for t := 0; t < step.Len(); t++ {
							detailStep := step.Index(t).Interface()
							if reflect.ValueOf(detailStep).Kind() == reflect.Struct {
								vOfDetailStep := reflect.ValueOf(detailStep)
								if vOfDetailStep.NumField() == 2 {
									switch fmt.Sprintf("%v", vOfDetailStep.Field(0)) {
									case "webDriver":
										webDriver = vOfDetailStep.Field(1)

									case "description":
										description = vOfDetailStep.Field(1)

									case "webElement":
										webElement = vOfDetailStep.Field(1)

									case "actions":
										actions = vOfDetailStep.Field(1)

									case "timeout":
										timeout = vOfDetailStep.Field(1)

									case "tab":
										tab = vOfDetailStep.Field(1)

									case "script":
										script = vOfDetailStep.Field(1)

									case "newUrl":
										newUrl = vOfDetailStep.Field(1)
									case "checkElements":
										checkElements = vOfDetailStep.Field(1)

									default:
										log.Error("vOfDetailStep.Field " +
											fmt.Sprintf("%v", vOfDetailStep.Field(0)) + " Chưa được định nghía")
									}
								} else {
									log.Error("vOfDetailStep.NumField " + strconv.Itoa(vOfDetailStep.NumField()) +
										" Chưa được định nghía")
								}
							}
						}
					}
					switch fmt.Sprintf("%v", webDriver) {
					case "Title":
						plusInfoActionTesting := testing.PlusInfoActionTesting{
							Page:            fmt.Sprintf("%v", vHook.Field(0)),
							OrdinalStep:     0,
							DescriptionStep: fmt.Sprintf("%v", description),
							JobID:           jobID,
							TestId:          testId.String(),
							WebDriver:       fmt.Sprintf("%v", webDriver),
							Action:          "",
							Data:            "",
						}
						err := automate_plugin.WebDriverTitle(conctRemote, contentTesting, userSignIn,
							testingRepo, plusInfoActionTesting)
						if err != nil {
							automate_plugin.PluginScreenShotError(conctRemote, contentTesting, userSignIn,
								testingRepo, plusInfoActionTesting)
							log.Error(err.Error())
							return err
						}
						//End case "Title"

					case "FindElement":
						// parse WebElement
						var byOfWebElement, valueOfWebElement, ignoredOfWebElement string
						iWebElement := reflect.ValueOf(webElement.Interface())
						if iWebElement.Kind() == reflect.Slice {
							for t := 0; t < iWebElement.Len(); t++ {
								v := reflect.ValueOf(iWebElement.Index(t).Interface())
								if v.NumField() == 2 {
									switch fmt.Sprintf("%v", v.Field(0)) {
									case "by":
										byOfWebElement = fmt.Sprintf("%v", v.Field(1))
									case "value":
										valueOfWebElement = fmt.Sprintf("%v", v.Field(1))
									case "ignored":
										ignoredOfWebElement = fmt.Sprintf("%v", v.Field(1))
									default:
										log.Error("v.Field(0) " +
											fmt.Sprintf("%v", v.Field(0)) + " Chưa được định nghía")
									}
								}
							}
						}
						structWebElement := testing.WebElement{
							By:      byOfWebElement,
							Value:   valueOfWebElement,
							Ignored: ignoredOfWebElement,
						}

						plusInfoActionTesting := testing.PlusInfoActionTesting{
							Page:            fmt.Sprintf("%v", vHook.Field(0)),
							OrdinalStep:     0,
							DescriptionStep: fmt.Sprintf("%v", description),
							JobID:           jobID,
							TestId:          testId.String(),
							WebDriver:       fmt.Sprintf("%v", webDriver),
							Action:          "",
							Data:            "",
							WebElement:      structWebElement,
							Actions:         actions,
						}
						err := automate_plugin.PluginFindElement(conctRemote, contentTesting, userSignIn,
							testingRepo, plusInfoActionTesting)
						if err != nil {
							plusInfoActionTesting.Data1 = fmt.Sprintf("%v", err.Error())
							automate_plugin.PluginScreenShotError(conctRemote, contentTesting, userSignIn,
								testingRepo, plusInfoActionTesting)
							log.Error(err.Error())
							return err
						}
						//End case "FindElement"

					case "FindElements":
						// parse WebElement
						var byOfWebElement, valueOfWebElement, ignoredOfWebElement string
						iWebElement := reflect.ValueOf(webElement.Interface())
						if iWebElement.Kind() == reflect.Slice {
							for t := 0; t < iWebElement.Len(); t++ {
								v := reflect.ValueOf(iWebElement.Index(t).Interface())
								if v.NumField() == 2 {
									switch fmt.Sprintf("%v", v.Field(0)) {
									case "by":
										byOfWebElement = fmt.Sprintf("%v", v.Field(1))
									case "value":
										valueOfWebElement = fmt.Sprintf("%v", v.Field(1))
									case "ignored":
										ignoredOfWebElement = fmt.Sprintf("%v", v.Field(1))
									default:
										log.Error("v.Field(0) " +
											fmt.Sprintf("%v", v.Field(0)) + " Chưa được định nghía")
									}
								}
							}
						}
						structWebElement := testing.WebElement{
							By:      byOfWebElement,
							Value:   valueOfWebElement,
							Ignored: ignoredOfWebElement,
						}

						//parse CheckElements
						var countOfCheck string
						var memberOEOfCheck, declareE4ClickOfCheck reflect.Value
						iCheckElements := reflect.ValueOf(checkElements.Interface())
						if iWebElement.Kind() == reflect.Slice {
							for t := 0; t < iCheckElements.Len(); t++ {
								v := reflect.ValueOf(iCheckElements.Index(t).Interface())
								if v.NumField() == 2 {
									switch fmt.Sprintf("%v", v.Field(0)) {
									case "countElements":
										countOfCheck = fmt.Sprintf("%v", v.Field(1))
									case "memberOfElements":
										memberOEOfCheck = v.Field(1)
									case "declareElement4Click":
										declareE4ClickOfCheck = v.Field(1)
									default:
										log.Error("v.Field(0) " +
											fmt.Sprintf("%v", v.Field(0)) + " Chưa được định nghía")
									}
								}
							}
						}
						structCheckElements := testing.CheckElements{
							CountElements:        countOfCheck,
							MemberOfElements:     memberOEOfCheck,
							DeclareElement4Click: declareE4ClickOfCheck,
						}

						plusInfoActionTesting := testing.PlusInfoActionTesting{
							Page:            fmt.Sprintf("%v", vHook.Field(0)),
							OrdinalStep:     0,
							DescriptionStep: fmt.Sprintf("%v", description),
							JobID:           jobID,
							TestId:          testId.String(),
							WebDriver:       fmt.Sprintf("%v", webDriver),
							Action:          "",
							Data:            "",
							WebElement:      structWebElement,
							CheckElements:   structCheckElements,
							Actions:         actions,
						}
						err := automate_plugin.PluginFindElements(conctRemote, contentTesting, userSignIn,
							testingRepo, plusInfoActionTesting)
						if err != nil {
							plusInfoActionTesting.Data1 = fmt.Sprintf("%v", err.Error())
							automate_plugin.PluginScreenShotError(conctRemote, contentTesting, userSignIn,
								testingRepo, plusInfoActionTesting)
							log.Error(err.Error())
							return err
						}
						//End case "FindElement"

					case "SetImplicitWaitTimeout":
						plusInfoActionTesting := testing.PlusInfoActionTesting{
							Page:            fmt.Sprintf("%v", vHook.Field(0)),
							OrdinalStep:     0,
							DescriptionStep: fmt.Sprintf("%v", description),
							JobID:           jobID,
							TestId:          testId.String(),
							WebDriver:       fmt.Sprintf("%v", webDriver),
							Timeout:         fmt.Sprintf("%v", timeout),
							Action:          "",
							Data:            "",
						}
						err := automate_plugin.SetImplicitWaitTimeout(conctRemote, contentTesting, userSignIn,
							testingRepo, plusInfoActionTesting)
						if err != nil {
							plusInfoActionTesting.Data1 = fmt.Sprintf("%v", err.Error())
							automate_plugin.PluginScreenShotError(conctRemote, contentTesting, userSignIn,
								testingRepo, plusInfoActionTesting)
							log.Error(err.Error())
							return err
						}
						//End case "SetImplicitWaitTimeout"

					case "SwitchWindow":
						plusInfoActionTesting := testing.PlusInfoActionTesting{
							Page:            fmt.Sprintf("%v", vHook.Field(0)),
							OrdinalStep:     0,
							DescriptionStep: fmt.Sprintf("%v", description),
							JobID:           jobID,
							TestId:          testId.String(),
							WebDriver:       fmt.Sprintf("%v", webDriver),
							Tab:             fmt.Sprintf("%v", tab),
							Action:          "",
							Data:            "",
						}
						err := automate_plugin.SwitchWindow(conctRemote, contentTesting, userSignIn,
							testingRepo, plusInfoActionTesting)
						if err != nil {
							plusInfoActionTesting.Data1 = fmt.Sprintf("%v", err.Error())
							automate_plugin.PluginScreenShotError(conctRemote, contentTesting, userSignIn,
								testingRepo, plusInfoActionTesting)
							log.Error(err.Error())
							return err
						}
					//End case "SwitchWindow"

					case "ExecuteScript":
						plusInfoActionTesting := testing.PlusInfoActionTesting{
							Page:            fmt.Sprintf("%v", vHook.Field(0)),
							OrdinalStep:     0,
							DescriptionStep: fmt.Sprintf("%v", description),
							JobID:           jobID,
							TestId:          testId.String(),
							WebDriver:       fmt.Sprintf("%v", webDriver),
							Script:          fmt.Sprintf("%v", script),
							Action:          "",
							Data:            "",
						}
						err := automate_plugin.ExecuteScript(conctRemote, contentTesting, userSignIn,
							testingRepo, plusInfoActionTesting)
						if err != nil {
							plusInfoActionTesting.Data1 = fmt.Sprintf("%v", err.Error())
							automate_plugin.PluginScreenShotError(conctRemote, contentTesting, userSignIn,
								testingRepo, plusInfoActionTesting)
							log.Error(err.Error())
							return err
						}
					//End case "SwitchWindow"

					case "SleepAction":
						plusInfoActionTesting := testing.PlusInfoActionTesting{
							Page:            fmt.Sprintf("%v", vHook.Field(0)),
							OrdinalStep:     0,
							DescriptionStep: fmt.Sprintf("%v", description),
							JobID:           jobID,
							TestId:          testId.String(),
							WebDriver:       fmt.Sprintf("%v", webDriver),
							Timeout:         fmt.Sprintf("%v", timeout),
							Action:          "",
							Data:            "",
						}
						err := automate_plugin.SleepAction(conctRemote, contentTesting, userSignIn,
							testingRepo, plusInfoActionTesting)
						if err != nil {
							plusInfoActionTesting.Data1 = fmt.Sprintf("%v", err.Error())
							automate_plugin.PluginScreenShotError(conctRemote, contentTesting, userSignIn,
								testingRepo, plusInfoActionTesting)
							log.Error(err.Error())
							return err
						}
					//End case "SleepAction"

					case "Screenshot":
						plusInfoActionTesting := testing.PlusInfoActionTesting{
							Page:            fmt.Sprintf("%v", vHook.Field(0)),
							OrdinalStep:     0,
							DescriptionStep: fmt.Sprintf("%v", description),
							JobID:           jobID,
							TestId:          testId.String(),
							WebDriver:       fmt.Sprintf("%v", webDriver),
							Timeout:         fmt.Sprintf("%v", timeout),
							Action:          "",
							Data:            "",
						}
						err := automate_plugin.Screenshot(conctRemote, contentTesting, userSignIn,
							testingRepo, plusInfoActionTesting)
						if err != nil {
							plusInfoActionTesting.Data1 = fmt.Sprintf("%v", err.Error())
							automate_plugin.PluginScreenShotError(conctRemote, contentTesting, userSignIn,
								testingRepo, plusInfoActionTesting)
							log.Error(err.Error())
							return err
						}
					//End case "SleepAction"
					case "AccessNewUrl":
						plusInfoActionTesting := testing.PlusInfoActionTesting{
							Page:            fmt.Sprintf("%v", vHook.Field(0)),
							OrdinalStep:     0,
							DescriptionStep: fmt.Sprintf("%v", description),
							JobID:           jobID,
							TestId:          testId.String(),
							WebDriver:       fmt.Sprintf("%v", webDriver),
							NewUrl:          fmt.Sprintf("%v", newUrl),
							Action:          "",
							Data:            "",
						}
						err := automate_plugin.AccessNewUrl(conctRemote, contentTesting, userSignIn,
							testingRepo, plusInfoActionTesting)
						if err != nil {
							plusInfoActionTesting.Data1 = fmt.Sprintf("%v", err.Error())
							automate_plugin.PluginScreenShotError(conctRemote, contentTesting, userSignIn,
								testingRepo, plusInfoActionTesting)
							log.Error(err.Error())
							return err
						}

					default:
						log.Error("webDriver " + fmt.Sprintf("%v", webDriver) +
							" Chưa được định nghía")
					}
				}
			}
		} else {
			log.Error("reflect.ValueOf(hook).Kind() không phải kiểu Struct")
		}
	} //End for range contentTesting.Hooks
	return nil
}
