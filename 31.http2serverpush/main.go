package main

import (
	"html/template"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.SetHTMLTemplate(html)

	router.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			time.Sleep(2 * time.Second)
			// use pusher.Push() to do server push
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	// Listen and Server in https://127.0.0.1:8080
	router.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}

/*
1.下载windows版 openssl
https://blog.csdn.net/qq_39081974/article/details/81059022
http://slproweb.com/products/Win32OpenSSL.html
2.生成RSA私钥
$ mkdir testdata
$ openssl genrsa -out ./testdata/server.key 1024
3.生成数字证书
$ openssl req -new -x509 -key ./testdata/server.key -out ./testdata/server.pem -days 365
参考：https://github.com/gin-gonic/examples/tree/master/http2
如命令在powershell执行不了，到cmd下执行
*/
