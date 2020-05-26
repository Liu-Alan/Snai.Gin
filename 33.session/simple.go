package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 创建基于cookie的存储引擎，snai.alan.123 参数是用于加密的密钥
	store := cookie.NewStore([]byte("snai.alan.123"))
	// 设置session中间件，参数snai.session，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	router.Use(sessions.Sessions("snai.session", store))

	router.GET("/get", func(c *gin.Context) {
		//初始化session对象
		session := sessions.Default(c)

		//通过session.Get读取session值
		//session是键值对格式数据，因此需要通过key查询数据
		name := session.Get("name")
		if name != "alan" {
			//设置session数据
			session.Set("name", "alan")

			//保存session数据
			session.Save()

			c.JSON(http.StatusOK, gin.H{"msg": "set session success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": "session is " + name.(string)})
		}
	})

	router.GET("/delete", func(c *gin.Context) {
		//初始化session对象
		session := sessions.Default(c)

		//删除session
		session.Delete("name")

		//保存session数据
		session.Save()

		name := session.Get("name")
		if name != "alan" {
			c.JSON(http.StatusOK, gin.H{"msg": "delete session success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": "delete session fail"})
		}
	})

	router.GET("/clear", func(c *gin.Context) {
		//初始化session对象
		session := sessions.Default(c)

		//删除整个session
		session.Clear()

		//保存session数据
		session.Save()

		name := session.Get("name")
		if name != "alan" {
			c.JSON(http.StatusOK, gin.H{"msg": "clear session success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": "clear session fail"})
		}
	})

	router.Run()

}
