package router

import (
	"github.com/gin-gonic/gin"
	"github.com/5MofDream/apollo/moltencore"
	"github.com/5MofDream/apollo/app/controller"
)

func init() {
	server := moltencore.Moltencore().GinServer().Server()

	server.GET("/debug", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	hb := controller.ExampleController{}
	hb.DI()
	server.GET("/example", hb.Example)

}
