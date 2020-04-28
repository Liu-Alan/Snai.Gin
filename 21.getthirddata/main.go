package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/somedatafromreader", func(c *gin.Context) {
		response, err := http.Get("http://snaill.net/sitedata/image/nginx_1_2.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
		}
	})
}
