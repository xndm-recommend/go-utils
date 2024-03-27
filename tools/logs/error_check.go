package logs

/*
有关报错打印的封装
*/
import (
	"errors"
	"reflect"
	"runtime"
	"strconv"

	"github.com/cihub/seelog"
)

// 错误处理函数
func Debugs(err string, params ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	seelog.Debugf(file+":"+strconv.Itoa(line)+" "+err, params...)
}

// 错误处理函数info
func Infos(err string, params ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	seelog.Infof(file+":"+strconv.Itoa(line)+" "+err, params...)
}

// 错误处理函数Warnf
func Warns(field string, params ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	_ = seelog.Warnf(file+":"+strconv.Itoa(line)+" "+field, params...)
}

// 错误处理函数Error
func CommonErr(err error, params ...interface{}) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		ddd := seelog.Errorf("%s:%d: %v", file, line, err)
		println(ddd.Error())
	}
}

// 错误处理函数Critical
func FatalErr(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		_ = seelog.Criticalf("Important error:", file, ":", line, err)
		panic(err)
	}
}

func CheckEmptyValue(val interface{}) {
	if reflect.TypeOf(val).Kind() == reflect.Int {
		if val.(int) == 0 {
			panic("this value shouldn't be 0")
		}
	} else if reflect.TypeOf(val).Kind() == reflect.Int64 {
		if val.(int64) == 0 {
			panic(`this value shouldn't be 0`)
		}
	} else if reflect.TypeOf(val).Kind() == reflect.String {
		if val.(string) == "" {
			panic(`this value shouldn't be ""`)
		}
	} else if reflect.TypeOf(val).Kind() == reflect.Float32 {
		if val.(float32) == 0.0 {
			panic(`this value shouldn't be 0.0`)
		}
	} else if reflect.TypeOf(val).Kind() == reflect.Float64 {
		if val.(float64) == 0.0 {
			panic(`this value shouldn't be 0.0`)
		}
	} else if reflect.TypeOf(val).Kind() == reflect.Slice {
		if len(val.([]interface{})) == 0 {
			panic(`this value shouldn't be empty slice`)
		}
	} else if reflect.TypeOf(val).Kind() == reflect.Map {
		if len(val.(map[interface{}]interface{})) == 0 {
			panic(`this value shouldn't be empty map`)
		}
	}
}

func CheckValueStat(v, min, max int) int {
	if v > max {
		CommonErr(errors.New("input value is too large!!!"))
		return max
	}
	if v < min {
		CommonErr(errors.New("input value is too small!!!"))
		return min
	}
	return v
}
