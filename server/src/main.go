package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ret := gin.Default()
	ret.GET("/helloworld", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	ret.Run(":3000")
}
