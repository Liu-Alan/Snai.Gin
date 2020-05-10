package main

import (
	"github.com/gin-gonic/gin"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@snaill.net", "phone": "123433"},
	"austin": gin.H{"email": "austin@snaill.net", "phone": "666"},
	"lena":   gin.H{"email": "lena@snaill.net", "phone": "523443"},
}

func main() {
	router := gin.Default()
}
