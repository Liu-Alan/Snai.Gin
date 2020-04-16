package main

import (
	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" josn:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main(){
	router :=gin.Default()

	router.POST("/loginjson"
	)
}