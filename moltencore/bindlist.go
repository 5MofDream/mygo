package moltencore

import (
	"apollo/lib"
	"apollo/conf"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Binder interface {
	Bind(c *lib.Container) *lib.Container
}

type SysBinder struct {
}

func init() {

}

func (s *SysBinder) Bind(c *lib.Container) {
	bindMasterDB(c)
}

// xorm
func bindMasterDB(c *lib.Container) {
	abstractDB := "DB"
	dbNode := new(lib.BindNode)
	conf := (*c.Get("config").(*conf.Config)).(*conf.ConfigImp)
	httpDatabase, err := conf.GetMap("http_database")
	if err == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", httpDatabase["username"].(string), httpDatabase["password"].(string), httpDatabase["host"].(string), httpDatabase["port"].(string), httpDatabase["dbname"].(string))
		engine, err := gorm.Open(httpDatabase["driver"].(string), dsn)
		if err != nil {
			panic(fmt.Sprintf("connect http db err database: %v ,dsn :%v , err:%v", httpDatabase, dsn, err))
		}
		dbNode.Fill(abstractDB, engine, nil, true, false)
		c.Bind(abstractDB, dbNode)
	}
}
