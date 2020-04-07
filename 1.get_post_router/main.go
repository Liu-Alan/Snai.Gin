package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context) {
	c.String(http.StatusOK, "get请求")
}

func posting(c *gin.Context) {
	c.String(http.StatusOK, "post请求")
}

func putting(c *gin.Context) {
	c.String(http.StatusOK, "put请求")
}

func deleting(c *gin.Context) {
	c.String(http.StatusOK, "delet请求")
}

func patching(c *gin.Context) {
	c.String(http.StatusOK, "patch请求")
}

func head(c *gin.Context) {
	c.String(http.StatusOK, "head请求")
}

func options(c *gin.Context) {
	c.String(http.StatusOK, "options请求")
}

func main() {
	router := gin.Default()

	router.GET("/someget", getting)
	router.POST("/somepost", posting)
	router.PUT("/someput", putting)
	router.DELETE("/somedelete", deleting)
	router.PATCH("/somepatch", patching)
	router.HEAD("/somehead", head)
	router.OPTIONS("/someoptions", options)

	//默认启动的是 8080端口，也可以自己定义启动端口,router.Run(":9090")
	router.Run()
}
