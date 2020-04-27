package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/jsonp", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// 访问 http://localhost:8080/JSONP?callback=call
		// 将会输出:   call({foo:"bar"})
		c.JSONP(http.StatusOK, data)
	})

	router.Run(":8080")
}
