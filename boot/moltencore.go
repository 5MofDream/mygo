package boot

import (
	"apollo/conf"
	"sync"
)

//core struct
type moltenCore struct {
	appName    string
	conf       *conf.ConfigImp //default conf
	httpServer interface{}     // default gin
	rpcServer  interface{}     //default rpcx
	rpcClient  interface{}     //default rpcx
}

var once sync.Once

var mc *moltenCore

func init() {
	initMoltenCore()
}

//get
func Moltencore() *moltenCore {
	return mc
}

func initMoltenCore() {
	once.Do(func() {
		mc = new(moltenCore)
		mc.setConf()
	})
}

func (mc *moltenCore) setConf() {
	mc.conf = conf.Conf()
}

//
func Eruption() {

}
