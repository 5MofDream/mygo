package conf

import (
	"os"
	"fmt"
	"apollo/lib"
)

//config
type Config interface {
	Get(key string) string
	GetList(key string) []interface{}
	GetMap(key string) map[string]interface{}
}

type ConfigImp struct {
	filenam []string
}

var config *ConfigImp

func init() {
	basePath, err := os.Getwd()
	if (err != nil) {
		panic("get base path error")
	}
	confPath := basePath + "/conf"
	isPath  := lib.PathExists(confPath)
	if(isPath == false ){
		panic(fmt.Sprintf("get error conf path: %v" , confPath))
	}
	configFileList := lib.GetAllFileByPath(confPath)

	fmt.Println(basePath)
}

func getConfigFileList(){

}

func Conf() *ConfigImp {
	return config
}
