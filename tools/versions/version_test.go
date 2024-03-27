package versions

import (
	"fmt"
	"github.com/xndm-recommend/go-utils/tools/logs"
	"testing"

	"github.com/xndm-recommend/go-utils/config"
)

const (
	Config_path = "../config/test.yaml"
)

func TestMysqlDbInfo_QueryStruct(t *testing.T) {
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	logs.CommonErr(err)
	dbinfo := Version{}
	dbinfo.GeVersionFromConf(&c, "Version")
	fmt.Println(dbinfo.GetAlgoVersion())
}
