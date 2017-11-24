package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	v1 := r.Group("/v1")
	{
		v1.GET("/shorten/*url", shorten)
		v1.GET("/resolve/:id", resolve)
	}
	return r
}
