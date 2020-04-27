package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 可以自定义返回的json数据前缀
	//router.SecureJsonPrefix(")]}',\n")

	router.GET("/somejson", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// 将会输出:   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	router.Run(":8080")
}
