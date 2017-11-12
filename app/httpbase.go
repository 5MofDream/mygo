package app

import (
	"apollo/lib"
	"apollo/moltencore"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type HttpBase struct {
	container *lib.Container
}

func init() {

}

func (hb *HttpBase)DI(){
	hb.container = moltencore.Moltencore().GinServer().Container()
}

func (hb *HttpBase)Example( c *gin.Context){
	DB := hb.container.Get("DB").(*xorm.Engine)
	ret , err := DB.Query("select * from users limit 1")
	if err != nil{

	}
	c.JSON(200, gin.H{
		"data": ret,
	})
}