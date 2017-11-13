package service

import (
	"apollo/lib"
)

type Base struct {
	Container *lib.Container
}

func (hb *Base) DI(Container *lib.Container) {
	hb.Container = Container
}
