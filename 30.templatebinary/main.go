package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}

		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func main() {
	router := gin.New()

	template, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	router.SetHTMLTemplate(template)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/index.html", gin.H{
			"Foo": "World",
		})
	})

	router.Run(":8080")
}
