package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	defer router.Run()

	// router.GET("/", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "https://www.google.com/")
	// })

	// router.POST("/test", func(c *gin.Context) {
	// 	c.Redirect(http.StatusFound, "/foo")
	// })

	router.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})
}
