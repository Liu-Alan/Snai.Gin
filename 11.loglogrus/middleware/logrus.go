package middleware

import (
	"time"

	"Snai.Gin/11.loglogrus/logging"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	logName := "gin.txt"

	// 新建log实例
	logger := logrus.New()

	lfsHook, err := logging.LfsHookToFile(logName)

	if err != nil {
		logger.Errorf("config local file system for logger error: %v", err)
	}

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
