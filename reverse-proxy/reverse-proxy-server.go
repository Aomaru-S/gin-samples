package main

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func proxy(c *gin.Context, realUrl string) {
	request := c.Request
	web, _ := url.Parse(realUrl)
	request.URL.Scheme = web.Scheme
	request.URL.Host = web.Host
	request.URL.Path = web.Path

	transport := http.DefaultTransport
	response, _ := transport.RoundTrip(request)

	for k, vv := range response.Header {
		for _, v := range vv {
			c.Header(k, v)
		}
	}

	_, _ = bufio.NewReader(response.Body).WriteTo(c.Writer)
}

func setupProxy() *gin.Engine {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		proxy(c, "http://localhost:8081/test")
	})

	return r
}

func main() {
	r := setupProxy()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
