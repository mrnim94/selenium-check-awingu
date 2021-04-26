package helper

import (
	"bufio"
	"runtime"
	"encoding/base64"
	"io/ioutil"
	"os"
	"selenium-check-awingu/log"
)

func Base64ImageToString(nameImage string) (string, error) {
	var logPath string
	if runtime.GOOS == "windows" {
		logPath = "./log_files/"
	} else {
		logPath = "../../log_files/"
	}
	// Open file on disk.
	f, err := os.Open(logPath+"screenshots/" + nameImage + ".png")
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	return encoded, nil
}
