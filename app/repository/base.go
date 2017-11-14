package repository

import (
	"github.com/5MofDream/apollo/lib"
	"github.com/jinzhu/gorm"
)

type Base struct {
	Container *lib.Container
}

func (hb *Base) DI(Container *lib.Container) {
	hb.Container = Container
}

func (hb *Base)DB()*gorm.DB{
	return hb.Container.Get("DB").(*gorm.DB)
}
