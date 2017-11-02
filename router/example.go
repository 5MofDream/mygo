package router

import (
	"github.com/gin-gonic/gin"
	"apollo/moltencore"
	"fmt"
)

func init() {
	server := moltencore.Moltencore().GinServer().Server()
	fmt.Println("routeraaaa")
	server.GET("/debug", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
