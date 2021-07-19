package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	//r.LoadHTMLFiles(getHTMLFilePaths("./auto-detect-document/document-root/")...)
	//fmt.Println(getHTMLFilePaths("./auto-detect-document/document-root/") )

	//r.LoadHTMLFiles("./auto-detect-document/document-root/index.html")
	//r.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "/index.html", gin.H{})
	//})
	r.GET("/:path", func(c *gin.Context) {
		path := c.Param("path")
		r.LoadHTMLFiles("./auto-detect-document/document-root/" + path + ".html")
		c.HTML(http.StatusOK, "/"+path+".html", gin.H{})
	})
	r.GET("/second-document/:path", func(c *gin.Context) {
		path := c.Param("path")
		r.LoadHTMLFiles("./auto-detect-document/document-root/second-document/" + path + ".html")
		c.HTML(http.StatusOK, "second-document/"+path+".html", gin.H{})
	})
	//r.GET("/:path", handler)
	//r.GET("/second-document/:path", handler2)

	return r
}

func main() {
	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func handler(c *gin.Context) {
	path := c.Param("path")
	c.HTML(http.StatusOK, "/"+path, gin.H{})
}

func handler2(c *gin.Context) {
	path := c.Param("path")
	c.HTML(http.StatusOK, "second-document/"+path, gin.H{})
}

func getHTMLFilePaths(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, getHTMLFilePaths(filepath.Join(dir, file.Name()))...)
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	for i, v := range paths {
		if !strings.HasSuffix(v, ".html") {
			paths = append(paths[:i], paths[i+1:]...)
		}
	}

	return paths
}
