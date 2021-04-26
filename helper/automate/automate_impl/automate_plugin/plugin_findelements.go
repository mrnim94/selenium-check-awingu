package automate_plugin

import (
	"fmt"
	"github.com/tebeka/selenium"
	"math/rand"
	"reflect"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	"selenium-check-awingu/model/testing"
	"selenium-check-awingu/repository"
	"strconv"
	"time"
)

func PluginFindElements(conctRemote selenium.WebDriver, contentTesting testing.YamlTesting, userSignIn model.JobsUser,
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

	log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) + " " + pIAT.DescriptionStep)

	switch plusInfoActionTesting.WebElement.By {
	case "ByCSSSelector", "ByXPATH":
		var elemByCSSSelectors []selenium.WebElement
		if plusInfoActionTesting.WebElement.By == "ByCSSSelector" {
			elemByCSSSelectors, err = conctRemote.FindElements(selenium.ByCSSSelector, plusInfoActionTesting.WebElement.Value)
		} else {
			elemByCSSSelectors, err = conctRemote.FindElements(selenium.ByXPATH, plusInfoActionTesting.WebElement.Value)
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
			// check countElements
			if (pIAT.CheckElements.CountElements) != "" {
				if pIAT.CheckElements.CountElements == strconv.Itoa(len(elemByCSSSelectors)) {
					pIAT.Action = "Check Elements OK: " + strconv.Itoa(len(elemByCSSSelectors))
					err = RecordActionToDB(testingRepo,
						contentTesting, userSignIn,
						pIAT)
					if err != nil {
						log.Error(err.Error())
						return err
					}
					log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) + " " +
						pIAT.DescriptionStep + " " + "CountElements = Object trên web = " + strconv.Itoa(len(elemByCSSSelectors)))
				} else {
					log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) + " " +
						pIAT.DescriptionStep + " " + "CountElements là " + pIAT.CheckElements.CountElements + " khác Object trên web = " +
						strconv.Itoa(len(elemByCSSSelectors)))
					pIAT.Data1 = "CountElements = " + pIAT.CheckElements.CountElements + " khác Object trên web = " +
						strconv.Itoa(len(elemByCSSSelectors))
					PluginScreenShotWarning(conctRemote, contentTesting, userSignIn,
						testingRepo, pIAT)
					//return nil
				}
			}

			//check memberOfElements
			if pIAT.CheckElements.MemberOfElements.IsValid() == true {
				slMemberOfElements := reflect.ValueOf(pIAT.CheckElements.MemberOfElements.Interface())
				if slMemberOfElements.Kind() == reflect.Slice {
					for t := 0; t < slMemberOfElements.Len(); t++ {
						slMemberOfElement := reflect.ValueOf(slMemberOfElements.Index(t).Interface())
						if slMemberOfElement.Kind() == reflect.Slice {
							var valueOfMemberOfElements string
							for t := 0; t < slMemberOfElement.Len(); t++ {
								dem := 0
								MemberOfElement := reflect.ValueOf(slMemberOfElement.Index(t).Interface())
								if MemberOfElement.NumField() == 2 {
									switch fmt.Sprintf("%v", MemberOfElement.Field(0)) {
									case "value":
										valueOfMemberOfElements = fmt.Sprintf("%v", MemberOfElement.Field(1))

									default:
										log.Error("vAction " + fmt.Sprintf("%v", MemberOfElement) +
											" Chưa được định nghía")
									}
								}
								for _, elemByCSSSelector := range elemByCSSSelectors {
									contentElement, err := elemByCSSSelector.Text()
									if err != nil {
										log.Error(err.Error())
										return err
									} else {
										if valueOfMemberOfElements == contentElement {
											dem++
										}
									}
								}
								if dem != 0 {
									pIAT = plusInfoActionTesting
									pIAT.Action = "Check Member " + valueOfMemberOfElements + " Of Elements"
									err := RecordActionToDB(testingRepo,
										contentTesting, userSignIn,
										pIAT)
									if err != nil {
										log.Error(err.Error())
										return err
									}
									log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) + " " +
										pIAT.DescriptionStep + " " + "memberOfElements " + valueOfMemberOfElements + " xuất hiện " + strconv.Itoa(dem) + " trên web")
								} else {
									log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) + " " +
										pIAT.DescriptionStep + " " + "memberOfElements " + valueOfMemberOfElements + " không xuất hiện trên web")
									pIAT.Data1 = "memberOfElements " + valueOfMemberOfElements + " không xuất hiện trên web"
									PluginScreenShotWarning(conctRemote, contentTesting, userSignIn,
										testingRepo, pIAT)
								}
							}
						}
					}
				} else {
					fmt.Println("day khong phai la slice")
				}
			}

			// Chạy Action
			if len(elemByCSSSelectors) > 0 {
				//if pIAT.CheckElements.DeclareElement4Click == "random" {
				//	min := 0
				//	max := len(elemByCSSSelectors) - 1
				//	ramdomIndex := rand.Intn(max-min) + min
				//
				//	for i, elemByCSSSelector := range elemByCSSSelectors {
				//		if i == ramdomIndex {
				//			err := ActionOnElement(elemByCSSSelector, contentTesting, userSignIn,
				//				testingRepo, plusInfoActionTesting)
				//			if err != nil {
				//				log.Error(err.Error())
				//				return err
				//			}
				//		}
				//	}
				//} else {
				//	pIAT = plusInfoActionTesting
				//	pIAT.Data1 = " DEV chỉ mới lập trình declareElement4Click: random"
				//	PluginScreenShotWarning(conctRemote, contentTesting, userSignIn,
				//		testingRepo, pIAT)
				//	return nil
				//}
				if pIAT.CheckElements.DeclareElement4Click.IsValid() == true {
					var valueOfStyle, valueOfRange string
					slDeclareElement4Clicks := reflect.ValueOf(pIAT.CheckElements.DeclareElement4Click.Interface())
					if slDeclareElement4Clicks.Kind() == reflect.Slice {
						for t := 0; t < slDeclareElement4Clicks.Len(); t++ {
							slDeclareElement4Click := slDeclareElement4Clicks.Index(t).Interface()
							if reflect.ValueOf(slDeclareElement4Click).Kind() == reflect.Struct {
								vDeclareElement := reflect.ValueOf(slDeclareElement4Click)
								if vDeclareElement.NumField() == 2 {
									switch fmt.Sprintf("%v", vDeclareElement.Field(0)) {
									case "style":
										valueOfStyle = fmt.Sprintf("%v", vDeclareElement.Field(1))
									case "range":
										valueOfRange = fmt.Sprintf("%v", vDeclareElement.Field(1))
									default:
										log.Error("vDeclareElement " + fmt.Sprintf("%v", vDeclareElement) +
											" Chưa được định nghía")
									}
								}
							}
						}
					}
					//nhận được giá trị style: random và range là nhiêu
					if valueOfStyle == "random" {
						var max int
						if valueOfRange == "" {
							max = len(elemByCSSSelectors) - 1
						}else {
							intRange, err := strconv.Atoi(valueOfRange)
							if err != nil {
								log.Error(err.Error())
								return err
							}
							max = intRange - 1
						}

						min := 0
						rand.Seed(time.Now().UnixNano())
						ramdomIndex := rand.Intn(max-min) + min

						for i, elemByCSSSelector := range elemByCSSSelectors {
							if i == ramdomIndex {
								err := ActionOnElement(elemByCSSSelector, contentTesting, userSignIn,
									testingRepo, plusInfoActionTesting)
								if err != nil {
									log.Error(err.Error())
									return err
								}
							}
						}
					}else {
						pIAT = plusInfoActionTesting
						pIAT.Data1 = " DEV chỉ mới lập trình declareElement4Click: random"
						PluginScreenShotWarning(conctRemote, contentTesting, userSignIn,
							testingRepo, pIAT)
						return nil
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
	//////////END log action
	return nil
}

func ActionOnElement(elemByCSSSelector selenium.WebElement, contentTesting testing.YamlTesting, userSignIn model.JobsUser,
	testingRepo repository.TestingRepo, plusInfoActionTesting testing.PlusInfoActionTesting) error {

	pIAT := plusInfoActionTesting

	log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) + " " +
		pIAT.DescriptionStep + " và thực hiện actions")
	iActions := reflect.ValueOf(pIAT.Actions.Interface())
	if iActions.Kind() == reflect.Slice {
		for t := 0; t < iActions.Len(); t++ {
			vActions := reflect.ValueOf(iActions.Index(t).Interface())
			if vActions.Kind() == reflect.Slice {
				var actionOfWebDriver, valueOfAction string
				for t := 0; t < vActions.Len(); t++ {
					vAction := reflect.ValueOf(vActions.Index(t).Interface())
					if vAction.NumField() == 2 {
						switch fmt.Sprintf("%v", vAction.Field(0)) {
						case "action":
							actionOfWebDriver = fmt.Sprintf("%v", vAction.Field(1))
						case "value":
							valueOfAction = fmt.Sprintf("%v", vAction.Field(1))
						default:
							log.Error("vAction " + fmt.Sprintf("%v", vAction) +
								" Chưa được định nghía")
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
						log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) +
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
						log.Info(contentTesting.NameTest + " - step " + strconv.Itoa(pIAT.OrdinalStep) +
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
							log.Info(contentTesting.NameTest+" - step "+strconv.Itoa(pIAT.OrdinalStep)+
								" Text mà tool lấy được: ", textOfElement)
						}
					default:
						log.Error("action của chưa được định nghĩa")
					}
				}
			}
		}
	}
	return nil
}
