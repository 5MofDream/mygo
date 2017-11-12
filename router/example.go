package router

import (
	"github.com/gin-gonic/gin"
	"apollo/moltencore"
	"apollo/app"
)

func init() {
	server := moltencore.Moltencore().GinServer().Server()

	server.GET("/debug", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	hb := app.HttpBase{}
	hb.DI()
	server.GET("/base" , hb.Example)


}
