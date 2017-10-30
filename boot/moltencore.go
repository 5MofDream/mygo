package boot

import (
	"apollo/conf"
	"sync"
)

//core struct
type moltenCore struct {
	appName    string
	conf       *conf.Config //default conf yml config
	httpServer interface{}     // default gin
	rpcServer  interface{}     //default rpcx
	rpcClient  interface{}     //default rpcx
}

var once sync.Once

var mc *moltenCore

// init one mc
func init() {
	once.Do(func() {
		mc = new(moltenCore)
		mc.InitMonltenCore(conf.Conf())
	})
}

// init mc
func ( mc *moltenCore)InitMonltenCore(cc *conf.Config){
	mc.RegisterConf(cc)
}
//get
func Moltencore() *moltenCore {
	return mc
}

func (mc *moltenCore) RegisterConf(cc *conf.Config) {
	mc.conf = cc
}

//
func fire() {

}
