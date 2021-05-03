package automate_impl

import (
	"github.com/tebeka/selenium"
	"selenium-check-awingu/helper/automate"
	"selenium-check-awingu/log"
)

//var webDriver selenium.WebDriver

type Selenium struct {
	Browser       string
	ConnectServer string
}

func NewSelenium(s *Selenium) automate.Automate {
	return &Selenium{
		Browser:       s.Browser,
		ConnectServer: s.ConnectServer,
	}
}

func (s *Selenium) getRemote() (selenium.WebDriver, error) {
	caps := selenium.Capabilities(map[string]interface{}{"browserName": s.Browser})
	connect, err := selenium.NewRemote(caps, s.ConnectServer)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return connect, err
}
