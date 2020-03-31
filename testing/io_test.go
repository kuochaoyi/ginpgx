package testing

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/kuochaoyi/ginpgx/common"
	"gopkg.in/yaml.v3"
)

func TestGetYAML(t *testing.T) {
	// t.Fatal("not implemented")
	// type DatabaseInfo struct {
	// 	Dev  map[string]string
	// 	Test map[string]string
	// 	Prod map[string]string
	// }
	// f := "config/pgx.yml"
	// fileByte := common.GetFileByte(f)
	// var di DatabaseInfo
	// yaml.Unmarshal(fileByte, &di)
	// for k, v := range di.Dev {
	// 	time.Sleep(300 * time.Millisecond)
	// 	fmt.Println(k, v)
	// 	t.Log("t.Logf", "Nothing")
	// }

	type P struct {
		Msg map[string]string
	}

	var d P
	dir, _ := os.Getwd()
	fp := path.Join(dir, "../common/message.yml")
	fmt.Println(fp)
	fileData, _ := ioutil.ReadFile(fp)
	fmt.Println(fileData)
	yaml.Unmarshal(fileData, &d)
	fmt.Println(d)

	if b, ok := d.Msg["1000"]; ok {
		t.Logf(b)
	} else {
		t.Logf("nil")
	}
	// }
}

func TestGetFileByte(t *testing.T) {
	// t.Fatal("not implemented")
	// type DatabaseInfo struct {
	// 	Dev  map[string]string
	// 	Test map[string]string
	// 	Prod map[string]string
	// }
	// f := "config/pgx.yml"
	// fileByte := common.GetFileByte(f)
	// var di DatabaseInfo
	// yaml.Unmarshal(fileByte, &di)
	// for k, v := range di.Dev {
	// 	time.Sleep(300 * time.Millisecond)
	// 	fmt.Println(k, v)
	// 	t.Log("t.Logf", "Nothing")
	// }

	type P struct {
		Msg map[string]string
	}

	var a P

	bb := common.GetFileByte("../common/message.yml")

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
