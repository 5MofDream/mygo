package conf

import (
	"os"
	"fmt"
	"apollo/lib"
	"io/ioutil"
	"github.com/smallfish/simpleyaml"
)


const CONFIG_PATH = "conf"
//config
type Config interface {
	Get(key string) string
	GetList(key string) []interface{}
	GetMap(key string) map[string]interface{}
}

type ConfigImp struct {
	configNode map[string]interface{}
}

var config *ConfigImp

func init() {
	configFileList:= getConfigFileList()
	//read file
	for _ , filename := range configFileList{
		config.configNode[filename] = parseYmlFile(filename)
	}
}

func getConfigFileList() []string{
	basePath, err := os.Getwd()
	if err != nil {
		panic("get base path error")
	}
	confPath := basePath + "/"+ CONFIG_PATH
	isPath  := lib.PathExists(confPath)
	if(isPath == false ){
		panic(fmt.Sprintf("get error conf path: %v" , confPath))
	}
	configFileList ,err := lib.GetAllFileByPath(confPath)
	if err != nil{
		panic(fmt.Sprintf("get conf file error:%v" , err))
	}
	return configFileList
}

func parseYmlFile(filename string)*simpleyaml.Yaml{
	fileData , err := ioutil.ReadFile(filename)
	if err!= nil{
		panic(fmt.Sprintf("read conf file error :%v" ,err ))
	}
	yaml,err := simpleyaml.NewYaml(fileData)
	if err != nil {
		panic("parse")
	}
	return yaml
}


func Conf() *ConfigImp {
	return config
}
