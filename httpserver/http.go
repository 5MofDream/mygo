package httpserver

import (
	"sync"
	"github.com/gin-gonic/gin"
)

type HttpServer interface {
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
		apolloHttpServer.server = gin.Default()//todo use New() to replace after know

	})
}

func initGinHttpServer(){

}

func Server()*HttpServer{
	return &httpServer
}