package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/long_async", func(c *gin.Context) {
		// 创建要在goroutine中使用的副本
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// 这里使用你创建的副本
			log.Println("Done! in path" + cCp.Request.URL.Path)
		}()
	})

	router.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// 这里没有使用goroutine，所以不用使用副本
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	router.Run(":8080")
}
