package httpserver

import (
	"sync"
	"github.com/gin-gonic/gin"
	"github.com/5MofDream/apollo/lib"
)

type HttpServer interface {
	Run(addr ...string) (error)
}

//imp HttpServer
type HttpServerImp struct {
	server    *gin.Engine
	container *lib.Container
}

var once sync.Once

var httpServer HttpServer

var apolloHttpServer *HttpServerImp

func init() {
	InitServer()
}

func InitServer() {
	once.Do(func() {
		apolloHttpServer = &HttpServerImp{}
		apolloHttpServer.initGinHttpServer()
		httpServer = apolloHttpServer
	})
}

func Server() *HttpServer {
	return &httpServer
}

func (hs *HttpServerImp) initGinHttpServer() {
	hs.server = gin.Default() //todo use New() to replace after know
	hs.container = lib.ContainerInstance()

}

func (hs *HttpServerImp) Server() *gin.Engine {
	return hs.server
}

func (hs *HttpServerImp) Container() *lib.Container {
	return hs.container
}

func (hs *HttpServerImp) Bind(abstract string, node *lib.BindNode) bool {
	return hs.container.Bind(abstract, node)
}

func (hs *HttpServerImp) Run(addr ...string) error {
	return hs.server.Run(addr...)
}
