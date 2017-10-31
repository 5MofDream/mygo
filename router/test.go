package router

import (
	"apollo/boot"
	"github.com/gin-gonic/gin"
)

func init(){
	server := boot.Moltencore().GinServer().Server()
	server.GET("/debug" , func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}