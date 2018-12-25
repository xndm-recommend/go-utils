package HTTP

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/xndm-recommend/go-utils/conf_read"
	"github.com/xndm-recommend/go-utils/errors_"
)

type HttpInfo struct {
	HttpClient *http.Client
	Url        string
	Para       []string
	TimeOut    int
}

func createHttpConns(this *HttpInfo, sLogin *config.HttpYamklData) {
	if 0 == sLogin.Time_out {
		errors_.CheckFatalErr(errors.New("can't read http post timeout"))
	}
	this.HttpClient = &http.Client{Timeout: time.Duration(sLogin.Time_out) * time.Millisecond}
	this.Url = sLogin.Url
	for _, param := range sLogin.Para {
		this.Para = append(this.Para, param)
	}
	this.TimeOut = sLogin.Time_out
}

// 线上设置url参数
func (this *HttpInfo) SetUrlPara(values ...interface{}) string {
	var url_tmp string = this.Url
	u, err := url.Parse(url_tmp)
	errors_.CheckCommonErr(err)
	for i, val := range values {
		sVal, err := val.(string)
		if false == err {
			sVal = strconv.Itoa(val.(int))
		}
		q := u.Query()
		if len(this.Para) <= i {
			errors_.CheckCommonErr(errors.New("Set Url Para error"))
		}
		q.Set(this.Para[i], sVal)
		u.RawQuery = q.Encode()
	}
	return u.String()
}

func GetHttpConnFromConf(this *conf_read.ConfigEngine, SectionName string) *HttpInfo {
	HttpInfo := new(HttpInfo)
	sLogin := getHttpFromConf(this, SectionName)
	createHttpConns(HttpInfo, sLogin)
	return HttpInfo
}
