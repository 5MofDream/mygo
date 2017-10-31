package boot

import (
	"apollo/conf"
	"sync"
	"apollo/httpserver"
)

//core struct
type moltenCore struct {
	appName    string
	conf       *conf.Config //default conf yml config
	httpServer *httpserver.HttpServer  // default gin
	rpcServer  interface{}  //default rpcx
	rpcClient  interface{}  //default rpcx
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

//get
func Moltencore() *moltenCore {
	return mc
}

// init mc
func (mc *moltenCore) InitMonltenCore(cc *conf.Config) {
	mc.RegisterConf(cc)
}

func (mc *moltenCore) RegisterConf(cc *conf.Config) {
	mc.conf = cc
}

func (mc *moltenCore) YamlConf() (cc *conf.ConfigImp) {
	cci, ok := (*mc.conf).(*conf.ConfigImp)
	if !ok {
		panic("get yaml error")
	}
	return cci
}

//
func fire() {

}
