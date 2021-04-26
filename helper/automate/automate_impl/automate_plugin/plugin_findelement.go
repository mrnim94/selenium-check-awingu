package automate_plugin

import (
	"fmt"
	"github.com/tebeka/selenium"
	"reflect"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/testing"
	"selenium-check-awingu/repository"
)

//function tìm kiếm element
func PluginFindElement(conctRemote selenium.WebDriver, contentTesting testing.YamlTesting, userSignIn model.JobsUser,
	testingRepo repository.TestingRepo, plusInfoActionTesting testing.PlusInfoActionTesting) error {
	var err error
	pIAT := plusInfoActionTesting
	pIAT.Action = "Begin"

	err = RecordActionToDB(testingRepo,
		contentTesting, userSignIn,
		pIAT)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	//////////END log action

	log.Info(contentTesting.NameTest + " - step " + pIAT.TestId + " " + pIAT.DescriptionStep)
	switch plusInfoActionTesting.WebElement.By {
	case "ByCSSSelector", "ByXPATH":
		var elemByCSSSelector selenium.WebElement
		if plusInfoActionTesting.WebElement.By == "ByCSSSelector" {
			elemByCSSSelector, err = conctRemote.FindElement(selenium.ByCSSSelector, plusInfoActionTesting.WebElement.Value)
		} else {
			elemByCSSSelector, err = conctRemote.FindElement(selenium.ByXPATH, plusInfoActionTesting.WebElement.Value)
		}

		if err != nil {
			log.Error(err.Error())
			if plusInfoActionTesting.WebElement.Ignored == "true" {
				return err
			} else {
				plusInfoActionTesting.Data1 = fmt.Sprintf("%v", err.Error())
				PluginScreenShotWarning(conctRemote, contentTesting, userSignIn,
					testingRepo, plusInfoActionTesting)
				return nil
			}
		} else {
			log.Info(contentTesting.NameTest + " - step " + pIAT.TestId + " " +
				pIAT.DescriptionStep + " và thực hiện actions")
			iActions := reflect.ValueOf(pIAT.Actions.Interface())
			if iActions.Kind() == reflect.Slice {
				for t := 0; t < iActions.Len(); t++ {
					vActions := reflect.ValueOf(iActions.Index(t).Interface())
					if vActions.Kind() == reflect.Slice {
						var actionOfWebDriver, valueOfAction, checkOfAction string
						for t := 0; t < vActions.Len(); t++ {
							vAction := reflect.ValueOf(vActions.Index(t).Interface())
							if vAction.NumField() == 2 {
								switch fmt.Sprintf("%v", vAction.Field(0)) {
								case "action":
									actionOfWebDriver = fmt.Sprintf("%v", vAction.Field(1))
								case "value":
									valueOfAction = fmt.Sprintf("%v", vAction.Field(1))
								case "check":
									checkOfAction = fmt.Sprintf("%v", vAction.Field(1))
								default:
									log.Error("vAction " + fmt.Sprintf("%v", vAction) +
										" Chưa được định nghía")
								}
							}
						}
						switch actionOfWebDriver {
						case "Sendkeys":
							var stringKeys string
							switch valueOfAction {
							case "usernameInDB":
								stringKeys = userSignIn.Username
							case "passwordInDB":
								stringKeys = userSignIn.Password
							default:
								stringKeys = valueOfAction
							}
							elemByCSSSelector.SendKeys(stringKeys)
							pIAT := plusInfoActionTesting
							pIAT.Action = "SendKeys"

							err := RecordActionToDB(testingRepo,
								contentTesting, userSignIn,
								pIAT)
							if err != nil {
								log.Error(err.Error())
								return err
							}
							//////////END log action
							log.Info(contentTesting.NameTest + " - step " + pIAT.TestId +
								" đã SendKeys vào Element")
						case "Click":
							elemByCSSSelector.Click()
							pIAT := plusInfoActionTesting
							pIAT.Action = "Click"

							err := RecordActionToDB(testingRepo,
								contentTesting, userSignIn,
								pIAT)
							if err != nil {
								log.Error(err.Error())
								return err
							}
							//////////END log action
							log.Info(contentTesting.NameTest + " - step " + pIAT.TestId +
								" đã Click vào Element")
						case "Text":
							textOfElement, err := elemByCSSSelector.Text()
							if err != nil {
								log.Error(err.Error())
								log.Error("Lấy text thất bại")
								//return err
							} else {
								pIAT := plusInfoActionTesting
								pIAT.Action = "Text"
								pIAT.Data = textOfElement
								err := RecordActionToDB(testingRepo,
									contentTesting, userSignIn,
									pIAT)
								if err != nil {
									log.Error(err.Error())
									return err
								}
								//////////END log action
								log.Info(contentTesting.NameTest+" - step "+pIAT.TestId+
									" Text mà tool lấy được: ", textOfElement)

								if checkOfAction != "" {
									if checkOfAction != textOfElement {
										pIAT.Data1 = "Tool lấy được content là: " +textOfElement+
											" không giống " +checkOfAction
										PluginScreenShotWarning(conctRemote, contentTesting, userSignIn,
											testingRepo, pIAT)
										log.Error("Tool lấy được content là: " +textOfElement+
											" không giống " +checkOfAction)
									}else {
										pIAT.Action = "Tool lấy được content là: " +textOfElement+
											" giống " +checkOfAction
										err := RecordActionToDB(testingRepo,
											contentTesting, userSignIn,
											pIAT)
										if err != nil {
											log.Error(err.Error())
											return err
										}
										//////////END log action
										log.Info(contentTesting.NameTest+" - step "+pIAT.TestId+
											" Check content findElement ok : ", textOfElement)
									}
								}
							}
						default:
							log.Error("action của chưa được định nghĩa")
						}
					}
				}
			}
		}

	default:
		log.Error("byOfWebElement của chưa được định nghĩa")
	}
	pIAT = plusInfoActionTesting
	pIAT.Action = "End"

	err = RecordActionToDB(testingRepo,
		contentTesting, userSignIn,
		pIAT)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info(contentTesting.NameTest + " - step " + pIAT.TestId + " " +
		pIAT.DescriptionStep +
		" Hoàn thành")
	//////////END log action
	return nil
}
