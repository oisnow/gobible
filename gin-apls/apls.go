package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Print("pingpong")
		c.JSON(200, gin.H{
			"message": "pingpong",
		})

	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
