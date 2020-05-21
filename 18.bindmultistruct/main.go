package main

import (
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func ShouldBindHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	// This c.ShouldBind consumes c.Request.Body and it cannot be reused.
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// Always an error is occurred by this because c.Request.Body is EOF now.
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		c.String(http.StatusOK, `the body no should`)
	}
}

func ShouldBindBodyWithHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	// This reads c.Request.Body and stores the result into the context.
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// At this time, it reuses body stored in the context.
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
	} else {
		c.String(http.StatusOK, `the body no should`)
	}
}

func main() {
	router := gin.Default()

	router.POST("/shouldbind", ShouldBindHandler)
	router.POST("/shouldbindbody", ShouldBindBodyWithHandler)

	router.Run()
}
