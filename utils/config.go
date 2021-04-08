package utils

import (
	"io/ioutil"
	"log"
)

func IsConfigFileExist() bool {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.Name() == "config.yaml" {
			return true
		}
	}
	return false
}
