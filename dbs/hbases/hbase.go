/*
 go build -tags kerberos -o event_receiver event_receiver.go
*/
package hbases

import (
	"context"
	"errors"
	"io"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/dbs/hbases/gohbase"
	"github.com/xndm-recommend/go-utils/dbs/hbases/gohbase/hrpc"
	"github.com/xndm-recommend/go-utils/tools/errs"
)

type HBaseDbInfo struct {
	Zkquorum  string
	Option    string
	Namespace string
	TableName map[string]string
	_client   gohbase.Client
}

func (hb *HBaseDbInfo) ConnectHBase(account string, zkquorum string) {
	auth := gohbase.Auth("KERBEROS")
	user := gohbase.EffectiveUser(account)
	options := []gohbase.Option{auth, user}
	hb.Zkquorum = zkquorum
	hb._client = gohbase.NewClient(zkquorum, options...)
}

func (hb *HBaseDbInfo) connectHBase(db *config.HBaseDbData) {
	auth := gohbase.Auth("KERBEROS")
	user := gohbase.EffectiveUser(db.User)
	options := []gohbase.Option{auth, user}
	hb.Zkquorum = db.ZK
	hb.Namespace = db.Namespace
	hb.TableName = db.TableName
	hb._client = gohbase.NewClient(db.ZK, options...)
}

//通过hb.PutsByRowkeyVersion(table, rowkey, values, hrpc.Timestamp(timestamp))调用，其中timestamp是time.Time类型，options也可以是其他 func(hrpc.Call)的函数
func (hb *HBaseDbInfo) PutsByRowkeyVersion(table, rowKey string, values map[string]map[string][]byte, options func(hrpc.Call) error) (err error) {
	putRequest, err := hrpc.NewPutStr(context.Background(), table, rowKey, values, options)
	errs.CheckCommonErr(err)
	_, err = hb._client.Put(putRequest)
	errs.CheckCommonErr(err)

	return
}

//指定表，通过options筛选数据，例如Families函数，或者filter函数
func (hb *HBaseDbInfo) GetsByOption(table, rowkey string, options ...func(hrpc.Call) error) (*hrpc.Result, error) {
	getRequest, err := hrpc.NewGetStr(context.Background(), table, rowkey, options...)
	if nil != err {
		errs.CheckCommonErr(err)
		return nil, err
	}
	res, err := hb._client.Get(getRequest)
	if nil != err {
		errs.CheckCommonErr(err)
		return nil, err
	}
	return res, err
}

//指定表，通过options筛选数据，例如Families函数，或者filter函数
func (hb *HBaseDbInfo) GetsByScanOption(table string, options ...func(hrpc.Call) error) (rsp []*hrpc.Result, err error) {
	var (
		scanRequest *hrpc.Scan
		res         *hrpc.Result
	)
	scanRequest, err = hrpc.NewScanStr(context.Background(), table, options...)
	if nil != err {
		errs.CheckCommonErr(err)
		return nil, err
	}
	scanner := hb._client.Scan(scanRequest)
	for {
		res, err = scanner.Next()
		if err == io.EOF || res == nil {
			break
		}
		if err != nil {
			errs.CheckCommonErr(err)
		}
		rsp = append(rsp, res)
	}
	return
}

func (this *HBaseDbInfo) Ping() error {
	if this._client != nil {
		return nil
	}
	return errors.New("连接为空")
}

func (this *HBaseDbInfo) GetClient() gohbase.Client {
	return this._client
}

func (this *HBaseDbInfo) GetDbConnFromConf(c *config.ConfigEngine, name string) {
	this.connectHBase(c.GetHBaseFromConf(name))
}