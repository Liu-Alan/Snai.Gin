package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "get",
		})
	})
	/*
		http.ListenAndServe(":8080", router)
	*/
	ser := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	ser.ListenAndServe()
}
