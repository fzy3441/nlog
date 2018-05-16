package nlog

import (
	"fmt"
	"github.com/ngaut/log"
	"runtime"
	"strings"
)

func SetOutputByName(path string) error {
	return log.SetOutputByName(path)
}

func Info(v ...interface{}) {
	log.Info(v...)
}

func Infof(format string, v ...interface{}) {
	log.Infof(format, v...)
}

func Debug(v ...interface{}) {
	log.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

func Warn(v ...interface{}) {
	log.Warning(v...)
}

func Warnf(format string, v ...interface{}) {
	log.Warningf(format, v...)
}

func Warning(v ...interface{}) {
	log.Warning(v...)
}

func Warningf(format string, v ...interface{}) {
	log.Warningf(format, v...)
}

// 日志
func Error(v ...interface{}) {
	v = append(v, source())
	log.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	v = append(v, source())
	log.Errorf(format+"%s", v...)

}

func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func source() string {
	var message string
	for i := 0; i <= 9; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok {
			pos := strings.LastIndex(file, "/")
			pos += 1
			message += fmt.Sprintf("\n%10s %-20s\tline:%-d", "file:", file[pos:], line)
		}
	}
	return message
}
