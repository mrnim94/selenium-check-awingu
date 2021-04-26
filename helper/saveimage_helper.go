package helper

import (
	"bytes"
	"image"
	"image/png"
	"runtime"
	"os"
	"selenium-check-awingu/log"
)

func HelpSaveImage(photo []byte, name string) error {
	img, _, err := image.Decode(bytes.NewReader(photo))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	var logPath string
	if runtime.GOOS == "windows" {
		logPath = "./log_files/"
	} else {
		logPath = "../../log_files/"
	}

	out, err := os.Create(logPath+"screenshots/" + name + ".png")
	if err != nil {
		log.Error(err.Error())
		return err
	}

	err = png.Encode(out, img)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
