package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Msg struct {
	Name    string `json:"user"`
	Message string `json:"message"`
	Number  int    `json:"number"`
}

func main() {
	router := gin.Default()
	defer router.Run()

	//http://localhost:8080/moreJSON?field_a=hello&field_b=world
	router.POST("/moreJSON", func(ctx *gin.Context) {
		var jsonMsg Msg
		ctx.Bind(&jsonMsg)
		ctx.JSON(http.StatusOK, gin.H{
			"user":    jsonMsg.Name,
			"message": jsonMsg.Message,
			"number":  jsonMsg.Number,
		})
	})
}
