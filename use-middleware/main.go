package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Execute\n")
	})

	return r
}

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "before\n")
		c.Next()
		c.String(http.StatusOK, "after")
	}
}

func main() {
	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
