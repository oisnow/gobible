package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("test")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello vscode")
	})
	r.Run(":8888")
}
