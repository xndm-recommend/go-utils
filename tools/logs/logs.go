package logs

import (
	"strings"

	"github.com/sirupsen/logrus"
)

//New 初始化log的配置
func New() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	return log
}

//Logger 默认的logger
var Logger = New()

//默认的Log对象
var defaultlog *logrus.Entry

//SetLog 设置log行为
func setLog(loglevel string, defaultField map[string]interface{}) *logrus.Entry {
	Logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "event",
			logrus.FieldKeyFunc:  "caller",
		}})
	upperLoglevel := strings.ToUpper(loglevel)
	switch {
	case upperLoglevel == "TRACE":
		Logger.SetLevel(logrus.TraceLevel)
	case upperLoglevel == "DEBUG":
		Logger.SetLevel(logrus.DebugLevel)
	case upperLoglevel == "INFO":
		Logger.SetLevel(logrus.InfoLevel)
	case upperLoglevel == "WARN":
		Logger.SetLevel(logrus.WarnLevel)
	case upperLoglevel == "ERROR":
		Logger.SetLevel(logrus.ErrorLevel)
	}
	log := Logger.WithFields(defaultField)
	return log
}

//Init 初始化默认的log对象
func Init(loglevel string, defaultField map[string]interface{}) {
	defaultlog = setLog(loglevel, defaultField)
}

//Trace 默认log打印Trace级别信息
func Trace(fields map[string]interface{}, args ...interface{}) {
	if defaultlog == nil {
		if fields == nil {
			Logger.Trace(args...)
		} else {
			Logger.WithFields(fields).Trace(args...)
		}
	} else {
		if fields == nil {
			defaultlog.Trace(args...)
		} else {
			defaultlog.WithFields(fields).Trace(args...)
		}
	}

}

//Debug 默认log打印Debug级别信息
func Debug(fields map[string]interface{}, args ...interface{}) {
	if defaultlog == nil {
		if fields == nil {
			Logger.Debug(args...)
		} else {
			Logger.WithFields(fields).Debug(args...)
		}
	} else {
		if fields == nil {
			defaultlog.Debug(args...)
		} else {
			defaultlog.WithFields(fields).Debug(args...)
		}
	}
}

//Info 默认log打印Info级别信息
func Info(fields map[string]interface{}, args ...interface{}) {
	if defaultlog == nil {
		if fields == nil {
			Logger.Info(args...)
		} else {
			Logger.WithFields(fields).Info(args...)
		}
	} else {
		if fields == nil {
			defaultlog.Info(args...)
		} else {
			defaultlog.WithFields(fields).Info(args...)
		}
	}
}

//Warn 默认log打印Warn级别信息
func Warn(fields map[string]interface{}, args ...interface{}) {
	if defaultlog == nil {
		if fields == nil {
			Logger.Warn(args...)
		} else {
			Logger.WithFields(fields).Warn(args...)
		}
	} else {
		if fields == nil {
			defaultlog.Warn(args...)
		} else {
			defaultlog.WithFields(fields).Warn(args...)
		}
	}
}

//Error 默认log打印Error级别信息
func Error(fields map[string]interface{}, args ...interface{}) {
	if defaultlog == nil {
		if fields == nil {
			Logger.Error(args...)
		} else {
			Logger.WithFields(fields).Error(args...)
		}
	} else {
		if fields == nil {
			defaultlog.Error(args...)
		} else {
			defaultlog.WithFields(fields).Error(args...)
		}
	}

}

//Fatal 默认log打印Fatal级别信息
func Fatal(fields map[string]interface{}, args ...interface{}) {
	if defaultlog == nil {
		if fields == nil {
			Logger.Fatal(args...)
		} else {
			Logger.WithFields(fields).Fatal(args...)
		}
	} else {
		if fields == nil {
			defaultlog.Fatal(args...)
		} else {
			defaultlog.WithFields(fields).Fatal(args...)
		}
	}
}

//Panic 默认log打印Panic级别信息
func Panic(fields map[string]interface{}, args ...interface{}) {
	if defaultlog == nil {
		if fields == nil {
			Logger.Panic(args...)
		} else {
			Logger.WithFields(fields).Panic(args...)
		}
	} else {
		if fields == nil {
			defaultlog.Panic(args...)
		} else {
			defaultlog.WithFields(fields).Panic(args...)
		}
	}
}
