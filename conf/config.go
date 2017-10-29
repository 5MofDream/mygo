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
	Get(key string) string
	GetList(key string) []interface{}
	GetMap(key string) map[string]interface{}
}

type ConfigImp struct {
	configNode interface{}
}

var apolloConfig *ConfigImp

var once sync.Once

func Conf() *ConfigImp {
	once.Do(func() {
		apolloConfig = &ConfigImp{}
		configFileList := getConfigFileList()
		//read file
		apolloConfig.configNode = parseYmlFile(configFileList)
	})
	return apolloConfig
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
	return ci.configNode.(*simpleyaml.Yaml)
}
