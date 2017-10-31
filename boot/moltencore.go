package boot

import (
	"apollo/conf"
	"sync"
	"apollo/httpserver"
	"apollo/lib"
)

//core struct
type moltenCore struct {
	appName    string
	conf       *conf.Config           //default conf yml config
	httpServer *httpserver.HttpServer // default gin
	rpcServer  interface{}            //default rpcx
	rpcClient  interface{}            //default rpcx
}

var once sync.Once

var mc *moltenCore

// init one mc
func init() {
	once.Do(func() {
		mc = new(moltenCore)
		mc.InitMonltenCore(conf.Conf(), httpserver.Server())
	})
}

//get
func Moltencore() *moltenCore {
	return mc
}

// init mc conf
func (mc *moltenCore) InitMonltenCore(cc *conf.Config, hs *httpserver.HttpServer) {
	mc.RegisterConf(cc)
	mc.RegisterHttp(hs)
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

//init mc http
func (mc *moltenCore) RegisterHttp(hs *httpserver.HttpServer) {
	mc.httpServer = hs
}

func (mc *moltenCore) GinServer() (*httpserver.HttpServerImp) {
	gs, ok := (*mc.httpServer).(*httpserver.HttpServerImp)
	if !ok {
		panic("get gin server error")
	}
	return gs
}

func (mc *moltenCore) HttpRun(addr ...string) error {
	return (*mc.httpServer).Run(addr...)
}

//start
func (mc *moltenCore) Fire() {
	var err error
	host, err := mc.YamlConf().Get("http_server_host")
	lib.PanicError(err)
	port, err := mc.YamlConf().Get("http_server_port")
	lib.PanicError(err)
	addr := host + ":" + port
	err = mc.HttpRun(addr)
	lib.PanicError(err)
}
