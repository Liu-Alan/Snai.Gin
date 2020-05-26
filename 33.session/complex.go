package main

import (
	"encoding/gob"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type user struct {
	Name string
	Age  int
}

func main() {
	//把user这个接头体注册进来，后面跨路由才可以获取到user数据
	gob.Register(user{})

	router := gin.Default()

	store := cookie.NewStore([]byte("snai.alan.123"))
	router.Use(sessions.Sessions("snai.session", store))

	router.GET("/set", func(c *gin.Context) {
		session := sessions.Default(c)

		//把结构体存入session中
		session.Set("user", user{"alan", 30})
		session.Save()

		c.JSON(http.StatusOK, gin.H{"msg": "set session success"})
	})

	router.GET("/get", func(c *gin.Context) {
		session := sessions.Default(c)

		user := session.Get("user")

		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	router.Run()
}
