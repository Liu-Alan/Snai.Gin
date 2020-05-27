package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router
}

func main() {
	router := PingRouter()

	router.Run("")
}
