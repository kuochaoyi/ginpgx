package common

import (
	"io/ioutil"
	"os"
	"path"
)

func GetFileToByte(filePath string) []byte {
	dir, _ := os.Getwd()
	fullPath := path.Join(dir, filePath)
	fileByte, _ := ioutil.ReadFile(fullPath)
	return fileByte
}
