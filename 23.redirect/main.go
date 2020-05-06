package main

import "github.com/gin-gonic/gin"

func main() {
	/*HTTP重定向
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})

	router.Run(":8080")
	*/
	/*Gin路由重定向*/
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})

	router.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	router.Run(":8080")
}
