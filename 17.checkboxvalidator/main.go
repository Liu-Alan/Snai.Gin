package main

import (
	"github.com/gin-gonic/gin"
)

type ColorsForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.Any("/colorsform", colorHandler)

	router.Run(":8080")
}

func colorHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		var colors ColorsForm

		if err := c.ShouldBind(&colors); err != nil {
			c.JSON(400, gin.H{
				"msg": err,
			})
			return
		}

		c.JSON(200, gin.H{"colors": colors.Colors})
	} else {
		c.HTML(200, "form.html", nil)
	}

}
