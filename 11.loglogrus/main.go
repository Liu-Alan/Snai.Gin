package main

import (
	"net/http"

	"Snai.Gin/11.loglogrus/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.LoggerToFile())

	router.GET("/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "get",
		})
	})

	router.POST("/post", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "post",
		})
	})

	router.Run(":8080")
}
