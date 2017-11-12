package lib

import (
	"sync"
)

// 容器结构
type Container struct {
	locker   *sync.Mutex
	bindList map[string]*BindNode
}

//register 注册时调用 ， boot为container完成所有register后初始化
type registerProvider interface {
	register()
	//boot()
}

func ContainerInstance() *Container {
	c := new(Container)
	c.locker = new(sync.Mutex)
	c.bindList = make(map[string]*BindNode)
	return c
}

//绑定
func (c *Container) Bind(abstract string, instance *BindNode) bool {
	c.locker.Lock()
	defer c.locker.Unlock()
	if _, ok := c.bindList[abstract]; ok == false {
		//fmt.Println(abstract , instance)
		//os.Exit(111)
		c.bindList[abstract] = instance
		return true
	} else {
		return false
	}
}

//获取节点,获取的时候调用provide ,延时加载
func (c *Container) Get(abstract string) (interface{}) {
	return c.bindList[abstract].value
}

// 绑定节点
type BindNode struct {
	abstract string
	value    interface{}
	provider registerProvider
	isSingle bool
	isDefer  bool
}

func (bn *BindNode) Fill(abstract string, value interface{}, provider registerProvider, isSingle bool, isDefer bool) {
	bn.abstract = abstract
	bn.isDefer = isDefer
	bn.value = value
	bn.provider = provider
	bn.isSingle = isSingle
}
