// command: go test -v testing/io_test.go
package testing

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/kuochaoyi/ginpgx/common"
)

func TestGetYAML(t *testing.T) {

	var msg common.Message
	dir, _ := os.Getwd()
	fp := path.Join(dir, "../common/message.yml")
	fmt.Println(fp)
	fileData, _ := ioutil.ReadFile(fp)
	fmt.Println(fileData)
	yaml.Unmarshal(fileData, &msg)
	fmt.Println(msg)

	if b, ok := msg.Msg[10000]; ok {
		t.Logf(b)
	} else {
		t.Logf("nil")
	}
}

/*
func TestGetFileToByte(t *testing.T) {
	t.Fatal("not implemented")
	f := "config/pgx.yml"
	yaml.Unmarshal(fileByte, &di)
	for k, v := range di.Dev {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(k, v)
		t.Log("t.Logf", "Nothing")
	}

	type P struct {
		Msg map[string]string
	}

	var a P

	bb := common.GetFileToByte("../common/message.yml")

	fmt.Println(bb)
	yaml.Unmarshal(bb, &a)
	fmt.Println(a)

	if b, ok := a.Msg["10000"]; ok {
		t.Logf(b)
		for k, v := range a.Msg {
			fmt.Printf("%s, %s ", k, v)
		}
	} else {
		t.Logf("nil")
	}
}
*/

func TestLoadDBinfo(t *testing.T) {

	var dbinfo common.DBinfo
	filebyte := common.GetFileToByte("../config/pgx.yml")
	yaml.Unmarshal(filebyte, &dbinfo)
	fmt.Println(dbinfo)
}
