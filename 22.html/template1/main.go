package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.LoadHTMLGlob("templates/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	router.Run(":8080")
}
