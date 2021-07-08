package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Message struct {
	Hello    string `json:"hello"`
	Response string `json:"response"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world", "response": "json"})
	})

	r.GET("/struct", func(c *gin.Context) {
		message := &Message{
			"world",
			"json",
		}
		c.JSON(http.StatusOK, message)
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
