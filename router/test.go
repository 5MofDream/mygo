package router

import (
	"apollo/boot"
	"github.com/gin-gonic/gin"
)

func init(){
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080
	server := boot.Moltencore().GinServer().Server()
	server.GET("/debug" , func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}