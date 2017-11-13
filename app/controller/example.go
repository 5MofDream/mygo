package controller

import (
	"apollo/app"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"apollo/app/service"
)

type ExampleController struct {
	app.HttpBase
}

func (hb *ExampleController) Example(c *gin.Context) {
	DB := hb.Container.Get("DB").(*gorm.DB)
	type Result struct {
		Id int
	}
	var result Result
	DB.Raw("select id from users limit 1").Scan(&result)
	myservcie := new(service.ExampleService)
	myservcie.DI(hb.Container)
	list := myservcie.DoExample()
	c.JSON(200, gin.H{
		"example": result,
		"users":list,
	})
}
