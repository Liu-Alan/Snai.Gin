package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	router.GET("/role/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}
