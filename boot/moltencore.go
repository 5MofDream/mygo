package boot

import (
	"apollo/conf"
	"sync"
	"apollo/httpserver"
	"apollo/lib"
	"fmt"
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
	//http server
	host, errHost := mc.YamlConf().Get("http_server_host")
	port, errPort := mc.YamlConf().Get("http_server_port")
	addr := host + ":" + port
	if(errHost == nil && errPort==nil){
		fmt.Println("start")
		go func(){
			fmt.Println("start")
			err = mc.HttpRun(addr)
			lib.PanicError(err)
			fmt.Println("end")
		}()
	}
	//rpc server

	//cli

	
	xaaa  := make(chan string)
	<- xaaa

}
