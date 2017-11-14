package rpcserver

import (
	"github.com/smallnest/rpcx/server"
	"sync"
	"github.com/5MofDream/apollo/lib"
)

type RpcServer interface {
	RegisterNode(name string, obj interface{}, metadata string) error
	Server(network string, addr string) error
}

type RpcServerImp struct {
	server *server.Server
	container *lib.Container
}

var once sync.Once

var rpcServer RpcServer

var apolloRpcServer *RpcServerImp

func init() {
	InitServer()
}

func InitServer() {
	once.Do(func() {
		apolloRpcServer = new(RpcServerImp)
		apolloRpcServer.initRpcxServer()
		rpcServer = apolloRpcServer
	})
}

//init rpcx
func (rs *RpcServerImp) initRpcxServer() {
	rs.server = new(server.Server)
}

func Instance() *RpcServer {
	return &rpcServer
}

func (rs *RpcServerImp) Server(network string, addr string) error {
	return rs.server.Serve(network, addr)
}

func (rs *RpcServerImp) RegisterNode(name string, obj interface{}, metadata string) error {
	return rs.server.RegisterName(name, obj, metadata)
}

func (rs *RpcServerImp) LoadContainer() error {

	return nil
}

func (rs *RpcServerImp) Bind(abstract string, node *lib.BindNode) bool {
	return rs.container.Bind(abstract, node)
}
func (rs *RpcServerImp) Container() *lib.Container {
	return rs.container
}