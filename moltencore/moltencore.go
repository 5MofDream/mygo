package moltencore

import (
	"github.com/5MofDream/apollo/conf"
	"sync"
	"github.com/5MofDream/apollo/httpserver"
	"github.com/5MofDream/apollo/lib"
	"github.com/5MofDream/apollo/rpcserver"
	"flag"
)

//core struct
type moltenCore struct {
	appName    string
	conf       *conf.Config           //default conf yml config
	httpServer *httpserver.HttpServer // default gin
	rpcServer  *rpcserver.RpcServer   //default rpcx
	container  *lib.Container         //base container
}

var once sync.Once

var mc *moltenCore

// init one mc
func init() {
	once.Do(func() {
		mc = new(moltenCore)
		mc.InitMonltenCore(conf.Conf(), httpserver.Server(), rpcserver.Instance())
	})
}

//get
func Moltencore() *moltenCore {
	return mc
}

// init mc conf
func (mc *moltenCore) InitMonltenCore(cc *conf.Config, hs *httpserver.HttpServer, rs *rpcserver.RpcServer) {
	mc.RegisterConf(cc)
	mc.RegisterHttp(hs)
	mc.RegisterRpcServer(rs)
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
	mc.bindHttpContainer()
}

func (mc *moltenCore) GinServer() (*httpserver.HttpServerImp) {
	gs, ok := (*mc.httpServer).(*httpserver.HttpServerImp)
	if !ok {
		panic("get gin server error")
	}
	return gs
}

//todo 抽象一个interface 暂时使用httpServerImp
func (mc *moltenCore) bindHttpContainer() {
	configNode := new(lib.BindNode)
	configNode.Fill("config", mc.conf, nil, true, false)
	bindRet := (*mc.httpServer).(*httpserver.HttpServerImp).Bind("config", configNode)
	if bindRet == false {
		panic("bind http container config error")
	}
	sysB := new(SysBinder)
	sysB.Bind((*mc.httpServer).(*httpserver.HttpServerImp).Container())
}

func (mc *moltenCore)bindBaseContainer(){
	configNode := new(lib.BindNode)
	configNode.Fill("config", mc.conf, nil, true, false)
	bindRet := mc.container.Bind("config", configNode)
	if bindRet == false {
		panic("bind base container config error")
	}
	sysBinderObj := new(SysBinder)
	sysBinderObj.Bind(mc.container)
}


func (mc *moltenCore)bindRpcContainer(){
	configNode := new(lib.BindNode)
	configNode.Fill("config", mc.conf, nil, true, false)
	bindRet := (*mc.rpcServer).(*rpcserver.RpcServerImp).Bind("config" , configNode)
	if bindRet == false {
		panic("bind http container config error")
	}
	sysBinderObj := new(SysBinder)
	sysBinderObj.Bind((*mc.rpcServer).(*rpcserver.RpcServerImp).Container())
}


func (mc *moltenCore) HttpRun(addr ...string) error {
	return (*mc.httpServer).Run(addr...)
}

//init mc rpc server
func (mc *moltenCore) RegisterRpcServer(rs *rpcserver.RpcServer) {
	mc.rpcServer = rs
}

func (mc *moltenCore) RpcxServer() (*rpcserver.RpcServerImp) {
	rs, ok := (*mc.rpcServer).(*rpcserver.RpcServerImp)
	if !ok {
		panic("get rpcx server error")
	}
	return rs
}

//todo set rpcx conf

func (mc *moltenCore) RpcRegister(name string, obj interface{}, metadata string) error {
	return (*mc.rpcServer).RegisterNode(name, obj, metadata)
}

func (mc *moltenCore) RpcServerRun(network string, addr string) error {
	return (*mc.rpcServer).Server(network, addr)
}

//cli flag
var ()
//start
func (mc *moltenCore) Fire() {

	flag.Parse()
	var err error
	//http server
	host, errHost := mc.YamlConf().Get("http_server_host")
	port, errPort := mc.YamlConf().Get("http_server_port")
	if errHost == nil && errPort == nil {
		addr := host + ":" + port
		//start http
		go func() {
			err = mc.HttpRun(addr)
			lib.PanicError(err)
		}()
	}
	//rpc server
	rpcNetwork, errRpcNetword := mc.YamlConf().Get("rpc_server_network")
	rpcAddr, errRpcAddr := mc.YamlConf().Get("rpc_server_addr")
	if errRpcNetword == nil && errRpcAddr == nil {
		go func() {
			mc.RpcServerRun(rpcNetwork, rpcAddr)
		}()
	}

	//cli

	xaaa := make(chan string)
	<-xaaa

}
