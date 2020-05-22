package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// MaxAge=0或未设置 表示未设置Max-Age属性，客户端会话期有效，关闭客户端删除
// MaxAge<0 表示立刻删除该cookie，等价于"Max-Age: 0"
// MaxAge>0 表示存在Max-Age属性，单位是秒

func main() {
	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "not set"
			c.SetCookie("gin_cookie", "test", 0, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run()
}
