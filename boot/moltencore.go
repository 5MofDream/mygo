package boot

import "apollo/conf"

//core struct
type moltenCore struct {
	appName string
	conf conf.Config  //default conf
	httpServer interface{} // default gin
	rpcServer interface{} //default rpcx
	rpcClient interface{} //default rpcx

}

var mc *moltenCore


func init(){
	initMoltenCore()
}

//get
func Moltencore() *moltenCore{
	return mc
}


func initMoltenCore(){
	mc := new(moltenCore)
	mc.initConf()
}

func ( mc *moltenCore) initConf(){

}

//
func Eruption(){

}



