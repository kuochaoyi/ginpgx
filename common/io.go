package common

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

func GetYAML(filePath string) interface{} {
	var d interface{}
	dir, _ := os.Getwd()
	fp := path.Join(dir, filePath)
	fileData, _ := ioutil.ReadFile(fp)
	yaml.Unmarshal(fileData, d)
	return d
}
