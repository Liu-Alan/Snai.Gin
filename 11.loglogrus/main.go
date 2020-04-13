package main

import (
	"net/http"

	"Snai.Gin/11.loglogrus/logging"
	"Snai.Gin/11.loglogrus/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	lfsHook, err := logging.LoggerToFile()

	if err != nil {
		logger.Errorf("config local file system for logger error: %v", err)
	}

	// 新增 Hook
	logger.AddHook(lfsHook)
}

func main() {
	router := gin.Default()
	router.Use(middleware.LoggerToFile())

	router.GET("/get", func(c *gin.Context) {
		logger.WithFields(logrus.Fields{
			"method": "get",
		}).Info("method get")

		c.JSON(http.StatusOK, gin.H{
			"msg": "get",
		})
	})

	router.POST("/post", func(c *gin.Context) {
		logger.WithFields(logrus.Fields{
			"method": "post",
		}).Info("method post")

		c.JSON(http.StatusOK, gin.H{
			"msg": "post",
		})
	})

	router.Run(":8080")
}
