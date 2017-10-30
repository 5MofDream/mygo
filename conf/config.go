package conf

import (
	"os"
	"fmt"
	"apollo/lib"
	"io/ioutil"
	"github.com/smallfish/simpleyaml"
	"sync"
)

//config path name
const CONFIG_PATH = "conf"

//config
type Config interface {
	Get(key string) (string, error)
	GetList(key string) ([]interface{}, error)
	GetMap(key string) (map[interface{}]interface{}, error)
}

type ConfigImp struct {
	configNode *simpleyaml.Yaml
}

var config Config

var apolloConfig *ConfigImp

var once sync.Once

func Conf() *Config {
	once.Do(func() {
		apolloConfig = &ConfigImp{}
		configFileList := getConfigFileList()
		apolloConfig.configNode = parseYmlFile(configFileList)
		config = apolloConfig
	})
	return &config
}

func getConfigFileList() []string {
	basePath, err := os.Getwd()
	if err != nil {
		panic("get base path error")
	}
	confPath := basePath + "/" + CONFIG_PATH
	isPath := lib.PathExists(confPath)
	if (isPath == false ) {
		panic(fmt.Sprintf("get error conf path: %v", confPath))
	}
	configFileList, err := lib.GetAllFileByPath(confPath)
	if err != nil {
		panic(fmt.Sprintf("get conf file error:%v", err))
	}
	return configFileList
}

func parseYmlFile(filenameList []string) *simpleyaml.Yaml {
	var fileData []byte
	for _, filename := range filenameList {
		tmpFileData, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(fmt.Sprintf("read conf file error :%v", err))
		}
		fileData = append(fileData, tmpFileData...)
	}

	yaml, err := simpleyaml.NewYaml(fileData)
	if err != nil {
		panic("parse")
	}
	return yaml
}

func (ci *ConfigImp) GetYmlNode() *simpleyaml.Yaml {
	return ci.configNode
}

func (ci *ConfigImp) Get(key string) (string, error) {
	return ci.GetYmlNode().Get(key).String()
}

func (ci *ConfigImp) GetList(key string) ([]interface{}, error) {
	return ci.GetYmlNode().Get(key).Array()
}

func (ci *ConfigImp) GetMap(key string) (map[interface{}]interface{}, error){
	return ci.GetYmlNode().Get(key).Map()
}
