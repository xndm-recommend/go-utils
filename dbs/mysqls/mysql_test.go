package mysqls_test

import (
	"github.com/xndm-recommend/go-utils/tools/logs"
	"testing"

	"github.com/xndm-recommend/go-utils/config"
	. "github.com/xndm-recommend/go-utils/dbs/mysqls"
)

const (
	Config_path = "../../config/test.yaml"
)

func TestMysqlDbInfo_QueryIdList(b *testing.T) {
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	logs.CommonErr(err)
	db := MysqlDbInfo{}
	db.GetDbConnFromConf(&c, "Comic_data")
	ids, _ := db.SelectStrList("select item_id from behavior limit 1")
	b.Log(ids)
}

func TestMysqlDbInfo_QueryStruct(t *testing.T) {
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	logs.CommonErr(err)
	dbinfo := MysqlDbInfo{}
	dbinfo.GetDbConnFromConf(&c, "Comic_data")

	//var cartoon_id1, cartoon_id2 int

	var names []string
	_ = dbinfo.GetDb().Select(&names, "SELECT cartoon_id FROM cartoon LIMIT 10")
	t.Log(names)
	t.Log(dbinfo.SelectStrList("select cartoon_id from cartoon limit ?", "10"))
}
