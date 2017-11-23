package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())

	v1 := router.Group("/v1")
	{
		v1.GET("/shorten/:url", shorten)
		v1.GET("/resolve/:hash", resolve)
	}
	
	router.Run(":8080")
}

func shorten(c *gin.Context)  {
	url := c.Param("url")
	c.String(http.StatusOK, "recieved request to shorten %s", url)
}

func resolve(c *gin.Context)  {
	hash := c.Param("hash")
	c.String(http.StatusOK, "recieved request to resolve %s", hash)
}
