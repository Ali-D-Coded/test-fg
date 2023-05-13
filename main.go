package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	
	api := gin.Default()

	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":"Hello World",
		})
	})
	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":"pong",
		})
	})

	api.Run(":8000")

}