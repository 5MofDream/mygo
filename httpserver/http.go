package httpserver

import (
	"sync"
	"github.com/gin-gonic/gin"
)

type HttpServer interface {
	Run(addr ...string)(error)
}

type HttpServerImp struct {
	server *gin.Engine
}

var once sync.Once

var httpServer HttpServer

var apolloHttpServer *HttpServerImp

func init() {
	InitServer()
}

func InitServer() {
	once.Do(func(){
		apolloHttpServer = & HttpServerImp{}
		apolloHttpServer.initGinHttpServer()
		httpServer = apolloHttpServer
	})
}



func Server()*HttpServer{
	return &httpServer
}

func (hs *HttpServerImp)initGinHttpServer(){
	apolloHttpServer.server = gin.Default()//todo use New() to replace after know
}

func (hs *HttpServerImp)Server()*gin.Engine{
	return hs.server
}

func (hs *HttpServerImp)Run(addr ...string)error{
	return hs.server.Run(addr...)
}
