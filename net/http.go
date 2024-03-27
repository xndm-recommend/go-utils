package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/xndm-recommend/go-utils/tools/logs"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/xndm-recommend/go-utils/config"
)

type HttpInfo struct {
	httpClient *http.Client
	url        string
	para       []string
	timeOut    int
}

func (this *HttpInfo) createHttpConns(sLogin *config.HttpData) {
	if 0 == sLogin.Time_out {
		logs.FatalErr(errors.New("can't read http post timeout"))
	}
	this.httpClient = &http.Client{Timeout: time.Duration(sLogin.Time_out) * time.Millisecond}
	this.url = sLogin.Url
	for _, param := range sLogin.Para {
		this.para = append(this.para, param)
	}
	this.timeOut = sLogin.Time_out
}

func (this *HttpInfo) GetHttpConnFromConf(c *config.ConfigEngine, name string) {
	this.createHttpConns(c.GetHttpFromConf(name))
}

// 线上设置url参数
func (this *HttpInfo) SetUrlPara(values ...interface{}) string {
	var url_tmp = this.url
	u, err := url.Parse(url_tmp)
	logs.CommonErr(err)
	for i, val := range values {
		sVal, err := val.(string)
		if false == err {
			sVal = strconv.Itoa(val.(int))
		}
		q := u.Query()
		if len(this.para) <= i {
			logs.CommonErr(errors.New("Set url para error"))
		}
		q.Set(this.para[i], sVal)
		u.RawQuery = q.Encode()
	}
	return u.String()
}

// 发送GET请求
// url:请求地址
// response:请求返回的内容
func (this *HttpInfo) HttpGet(url string) (response string, ok bool) {
	resp, err := this.httpClient.Get(url)
	if err != nil {
		logs.CommonErr(err)
		return "", false
	}
	defer resp.Body.Close()
	if 200 == resp.StatusCode {
		body, err := ioutil.ReadAll(resp.Body)
		logs.CommonErr(err)
		return string(body), true
	} else {
		return "", false
	}
}

// 发送GET请求
// url:请求地址
// response:请求返回的内容
func (this *HttpInfo) HttpGetBody(url string, body []byte) (response string, ok bool) {
	request, err := http.NewRequest("GET", url, bytes.NewReader(body))
	logs.CommonErr(err)
	resp, err := this.httpClient.Do(request)
	if err != nil {
		logs.CommonErr(err)
		return "", false
	}
	defer resp.Body.Close()
	if 200 == resp.StatusCode {
		body, err := ioutil.ReadAll(resp.Body)
		logs.CommonErr(err)
		return string(body), true
	} else {
		return "", false
	}
}

// 发送GET请求，失败后间隔一秒后重试，会重试指定次数
// url:请求地址
// times:重试次数
// response:请求返回的内容
func (this *HttpInfo) HttpGetDelayRetry(url string, times int) (response string, ok bool) {
	for i := 0; i < times; i++ {
		response, ok = this.HttpGet(url)
		if ok {
			break
		}
		time.Sleep(time.Second)
	}
	return response, ok
}

// 发送POST请求
// url:请求地址，types:POST请求提交的数据,contentType:请求体格式，如：application/json_struct
// content:请求放回的内容
func (this *HttpInfo) HttpPost(url string, data interface{}, contentType string) (content string, err error) {
	jsonStr, err := json.Marshal(data)
	logs.CommonErr(err)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	logs.CommonErr(err)
	req.Header.Set("Content-Type", contentType)
	defer req.Body.Close()
	resp, err := this.httpClient.Do(req)
	if err != nil {
		logs.CommonErr(err)
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	logs.CommonErr(err)
	return string(result), err
}
