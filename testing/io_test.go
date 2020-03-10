package testing

import (
	"fmt"
	"testing"

	"github.com/kuochaoyi/ginpgx/common"
)

type YAMLinfo struct {
	Info map[string]string
}

func TestGetYAML(t *testing.T) {
	// t.Fatal("not implemented")
	f := "config/pgx.yml"
	var y YAMLinfo
	y, _ = common.GetYAML(f).(YAMLinfo)
	for k, v := range y.Info {
		fmt.Println(k, v)
	}
}
