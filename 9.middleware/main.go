package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func BenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "BenchLogger")
		log.Println("BenchLogger")
		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Println(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "AuthRequired")
		log.Println("AuthRequired")
		// before request

		//c.Next()
		c.Abort()

		// after request
		latency := time.Since(t)
		log.Println(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	// 创建一个不包含中间件的路由器
	r := gin.New()

	// 全局中间件
	// 使用 Logger 中间件
	r.Use(gin.Logger())

	// 使用 Recovery 中间件
	r.Use(gin.Recovery())

	// 路由添加中间件，可以添加任意多个
	r.GET("/benchmark", BenchLogger(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "router middleware",
		})
	})

	// 路由组中添加中间件
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "login group middleware",
			})
		})

		authorized.POST("/submit", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "submit group middleware",
			})
		})

		authorized.POST("/read", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "read group middleware",
			})
		})

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "testing analytics group middleware",
			})
		})
	}

	r.Run(":8080")
}
