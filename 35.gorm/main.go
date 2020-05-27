package main

import (
	"net/http"
	"strconv"

	"Snai.Gin/35.gorm/dataaccess"
	"Snai.Gin/35.gorm/entities"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/adduser", func(c *gin.Context) {
		name := c.Query("name")
		address := c.Query("address")

		user := entities.User{
			Name:    name,
			Address: address,
		}

		err := dataaccess.AddUser(&user)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "增加成功",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
	})

	router.GET("/getuser", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))

		user, err := dataaccess.GetUser(id)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user": user,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
	})

	router.GET("/getusers", func(c *gin.Context) {
		users, err := dataaccess.GetUsers()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"users": users,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
	})

	router.GET("/modifyuser", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		name := c.Query("name")

		user, err := dataaccess.GetUser(id)
		if err == nil {
			user.Name = name
			err = dataaccess.ModifyUser(user)
			if err == nil {
				c.JSON(http.StatusOK, gin.H{
					"msg": "修改成功",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
	})

	router.GET("/deleteuser", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))

		user, err := dataaccess.GetUser(id)
		if err == nil {
			err = dataaccess.DeleteUser(user)
			if err == nil {
				c.JSON(http.StatusOK, gin.H{
					"msg": "删除成功",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
	})

	router.Run()
}
