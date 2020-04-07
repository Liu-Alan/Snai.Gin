package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "login",
			})
		})
		v1.POST("/submit", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "submit",
			})
		})
		v1.POST("/read", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "read",
			})
		})
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "login",
			})
		})
		v2.POST("/submit", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "submit",
			})
		})
		v2.POST("/read", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "read",
			})
		})
	}

	router.Run(":8080")
}
