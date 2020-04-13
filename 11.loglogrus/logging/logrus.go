package logging

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// log Hook设置
func LfsHookToFile(logName string) (logrus.Hook, error) {
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		logName+".%Y%m%d%H",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(logName),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(time.Hour*24),

		// 设置最大保存时间(30天)
		rotatelogs.WithMaxAge(time.Hour*24*30),
	)

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: logWriter,
		logrus.InfoLevel:  logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return lfsHook, err
}

func LoggerToFile() (logrus.Hook, error) {
	logName := "app.txt"

	//取Hook
	lfsHook, err := LfsHookToFile(logName)

	return lfsHook, err
}
