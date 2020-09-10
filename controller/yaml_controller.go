package controller

import (
	"io/ioutil"
	"log"
)

func GetYAMLFileContent(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
