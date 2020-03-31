package common

import (
	"io/ioutil"
	"os"
	"path"
)

func GetFileByte(filePath string) []byte {
	dir, _ := os.Getwd()
	fp := path.Join(dir, filePath)
	fileByte, _ := ioutil.ReadFile(fp)
	// yaml.Unmarshal(fileData, d)
	return fileByte
}
