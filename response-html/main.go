package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLFiles("./response-html/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Response-HTML",
			"message": "Hello, World! Response HTML.",
		})
	})

	return r
}

func main() {
	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
