package logs_test

import (
	"fmt"
	"github.com/xndm-recommend/go-utils/tools/logs"
	"testing"
)

const (
	logFile    = "../../config/test.yaml"
	logxmlFile = "seelog.xml"
)

//func TestErrLog1(b *testing.T) {
//	appConfig := &config.ConfigEngine{}
//	logs.LoggerSetup(logxmlFile)
//	logs.FatalErr(appConfig.Load(logFile))
//	logs.FatalErr(appConfig.SentryRavenInit("Sentry_dsn"))
//	//errs.CheckErrSendEmail(fmt.Errorf("asdfasdfsafdsdfs"))
//	raven.CaptureError(fmt.Errorf("asdfasdfsafdsdfs"), nil)
//}

//
//func TestErrLog(b *testing.T) {
//	appConfig := &config.ConfigEngine{}
//	logs.LoggerSetup(logxmlFile)
//	logs.FatalErr(appConfig.Load(logFile))
//	logs.FatalErr(appConfig.SentryRavenInit("Sentry_dsn"))
//	logs.CheckErrSendEmail(fmt.Errorf("asdfasdfsafdsdfs"))
//}
//
//func TestRusErrLog1(b *testing.T) {
//	appConfig := &config.ConfigEngine{}
//	logs.FatalErr(appConfig.Load(logFile))
//	logs.FatalErr(appConfig.SentryRavenInit("Sentry_dsn"))
//	logs.CheckLogrusCaptureError(fmt.Errorf("asdf"), nil, "yyy", "ppp")
//}

func TestCheckCommonErr(b *testing.T) {
	logs.LoggerSetup(logxmlFile)
	//seelog.ReplaceLogger(seelog.Disabled)
	//logger, _ := seelog.LoggerFromConfigAsFile("seelog.xml")
	//seelog.ReplaceLogger(logger)
	//defer seelog.Flush()
	//seelog.ReplaceLogger(seelog.Disabled)
	//logger, _ := seelog.LoggerFromConfigAsFile("seelog.xml")
	//seelog.ReplaceLogger(logger)
	//defer seelog.Flush()

	//appConfig := &config.ConfigEngine{}
	logs.CommonErr(fmt.Errorf("TestCheckCommonErr hahah"))
	//errs.FatalErr(appConfig.Load(logFile))
	//errs.FatalErr(appConfig.SentryRavenInit("Sentry_dsn"))
	//errs.CheckLogrusCaptureError(fmt.Errorf("asdf"), nil, "yyy", "ppp")
}
