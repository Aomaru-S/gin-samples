package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupWeb() *gin.Engine {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusOK, c.Request.RequestURI)
	})

	return r
}

func main() {
	r := setupWeb()
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
