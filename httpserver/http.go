package httpserver

import (
	"sync"
	"github.com/gin-gonic/gin"
)

type HttpServer interface {
	Run()(interface{})
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
		initGinHttpServer()
	})
}

func initGinHttpServer(){
	apolloHttpServer.server = gin.Default()//todo use New() to replace after know
}

func Server()*HttpServer{
	return &httpServer
}



func (hs *HttpServerImp)Server()*gin.Engine{
	return hs.server
}

func (hs *HttpServerImp)Run(addr ...string)error{
	return hs.server.Run(addr...)
}
