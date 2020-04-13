package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	logName := "gin.txt"

	logger := logrus.New()

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

	if err != nil {
		logger.Errorf("config local file system for logger error: %v", err)
	}

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

	// 新增 Hook
	logger.AddHook(lfsHook)

	// 写日志
	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()

		//处理请求
		c.Next()

		//结束时间
		endTime := time.Now()

		//执行时间
		latencyTime := endTime.Sub(startTime)

		//请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// post参数
		if c.Request.Method == "POST" {
			c.Request.ParseForm()
		}
		postParam := c.Request.PostForm.Encode()

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//客户端类型
		userAgent := c.Request.UserAgent()

		// 日志格式
		logger.WithFields(logrus.Fields{
			"statusCode":  statusCode,
			"latencyTime": latencyTime,
			"clientIP":    clientIP,
			"userAgent":   userAgent,
			"reqMethod":   reqMethod,
			"reqUri":      reqUri,
			"postParam":   postParam,
		}).Info()
	}
}
