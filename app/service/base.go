package service

import (
	"github.com/5MofDream/apollo/lib"
)

type Base struct {
	Container *lib.Container
}

func (hb *Base) DI(Container *lib.Container) {
	hb.Container = Container
}
