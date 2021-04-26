package testing

import (
	"gopkg.in/yaml.v2"
	"reflect"
)

type YamlTesting struct {
	Version      string            `yaml:"version"`
	NameTest     string            `yaml:"nameTest"`
	Browser      string            `yaml:"browser"`
	Web          string            `yaml:"web"`
	ResizeWindow map[string]string `yaml:"resizeWindow"`
	Actions      []string          `yaml:"actions"`
	Hooks        yaml.MapSlice     `yaml:"hooks"`
}

type PlusInfoActionTesting struct {
	Page            string
	OrdinalStep     int
	DescriptionStep string
	JobID           string
	TestId          string
	WebDriver       string
	Action          string
	Timeout         string
	Tab             string
	Script          string
	NewUrl          string
	Data            string
	Data1           string
	WebElement      WebElement
	CheckElements   CheckElements
	Actions         reflect.Value
}

type WebElement struct {
	By      string
	Value   string
	Ignored string
}

type CheckElements struct {
	CountElements        string
	MemberOfElements     reflect.Value
	DeclareElement4Click reflect.Value
}
