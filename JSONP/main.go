package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/JSONP", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		ctx.JSONP(http.StatusOK, data)
	})

	router.Run(":8085")
}
