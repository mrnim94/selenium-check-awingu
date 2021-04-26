package automate_impl

import (
	"fmt"
	"github.com/tebeka/selenium"
	"selenium-check-awingu/helper"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	"time"
)

func (s *Selenium) TestingConcurrencyAwingu(testing model.Testing) error {
	conctRemote, err := s.getRemote()
	if err != nil {
		log.Error(err.Error() + ">>>" + testing.Username)
		return err
	}
	//defer conctRemote.Quit()

	err = conctRemote.Get("https://sso.vngcloud.vn/cas/login?service=https://my.vngcloud.vn/sso%3Fr%3Dhttp%3A%2F%2Fmy.vngcloud.vn%2F%3Fhl%3Den")
	if err != nil {
		log.Error(err.Error())
		return err
	}

	conctRemote.SetImplicitWaitTimeout(time.Second * 60)

	if title, err := conctRemote.Title(); err == nil {
		fmt.Printf("Page title: %s\n", title)
	} else {
		log.Error(err.Error())
		return err
	}

	conctRemote.ResizeWindow("note", 500, 360)

	elemUserName, err := conctRemote.FindElement(selenium.ByCSSSelector, "input#username")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemUserName.SendKeys(testing.Username)
	log.Printf(testing.Username + "Đã điền được User")

	elemPassword, err := conctRemote.FindElement(selenium.ByCSSSelector, "input#password")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemPassword.SendKeys(testing.Password)
	log.Printf(testing.Username + "Đã điền được Password")

	//elemLogin, err := conctRemote.FindElement(selenium.ByCSSSelector, "input.btn.btn-block.btn-submit")
	//if err != nil {
	//	log.Error(err.Error())
	//	return err
	//}
	script := "document.querySelector('input.btn.btn-block.btn-submit').click()"
	args := []interface{}{}

	javaClick, err := conctRemote.ExecuteScript(script, args)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	fmt.Println(javaClick)
	//elemLogin.Click()
	//log.Printf(testing.Username + "Đã Nhan login")
	//
	//conctRemote.SetImplicitWaitTimeout(time.Second * 5)
	//
	//elemFirstAccept, err := conctRemote.FindElement(selenium.ByCSSSelector, "[qa-id='privacy-policy-accept-btn']")
	//if err != nil {
	//	log.Error(err.Error())
	//	log.Printf(testing.Username + "Đã Nhan Accept trước đó")
	//}else {
	//	elemFirstAccept.Click()
	//	log.Printf(testing.Username + "Đã Nhan Accept")
	//}
	//
	//elemGotIt, err := conctRemote.FindElement(selenium.ByCSSSelector, "[qa-id='tour-skip-btn']")
	//if err != nil {
	//	log.Error(err.Error())
	//	log.Printf(testing.Username + "Đã Nhan Got It trước đó")
	//}else {
	//	elemGotIt.Click()
	//	log.Printf(testing.Username + "Đã Nhan Got It")
	//}
	//
	//elemCheckLogin, err := conctRemote.FindElement(selenium.ByCSSSelector, "span.name.ng-binding")
	//if err != nil {
	//	log.Error(err.Error())
	//	log.Error(testing.Username + " login thất bại")
	//	return err
	//}
	//name, err := elemCheckLogin.Text()
	//if err != nil {
	//	log.Error(err.Error())
	//	log.Error(testing.Username + " login thất bại")
	//	return err
	//}
	//log.Printf("Login thành công", name)
	//
	//elemOpenWord, err := conctRemote.FindElement(selenium.ByXPATH,
	//	"//a[contains(text(),'Word')]")
	//if err != nil {
	//	log.Error(err.Error())
	//	return err
	//}
	//elemOpenWord.Click()
	//log.Printf(testing.Username + " đã mở Word")
	//
	//elemGotIt2, err := conctRemote.FindElement(selenium.ByCSSSelector, "[qa-id='tour-skip-btn']")
	//if err != nil {
	//	log.Error(err.Error())
	//	log.Printf(testing.Username + "Đã Nhan Got It trước đó")
	//}else {
	//	elemGotIt2.Click()
	//	log.Printf(testing.Username + "Đã Nhan Got It")
	//}

	shot, err := conctRemote.Screenshot()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	nameImage := testing.Username + "-testingccu-" + time.Now().Format("20060102150405")
	err = helper.HelpSaveImage(shot, nameImage)
	if err != nil {
		log.Error(err.Error())
		return err
	} else {
		log.Printf(testing.Username + "Đã lưu ảnh thành công " + nameImage)
	}

	elemTypeWordToApp, err := conctRemote.FindElement(selenium.ByCSSSelector, "#wsinput")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	for {
		elemTypeWordToApp.SendKeys("hello là xin chao")
		log.Printf(testing.Username + "Đã đánh văn bản")
		time.Sleep(time.Second * 2)
	}

	return nil
}
