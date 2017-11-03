package lib

import "sync"

// 容器结构
type Container struct {
	locker     *sync.Mutex
	bindList map[string]*BindNode
}

//register 注册时调用 ， boot为container完成所有register后初始化
type registerProvider interface {
	register()
	boot()
}

// 绑定节点
type BindNode struct {
	abstract string
	value    interface{}
	provider registerProvider
	isSingle bool
	isDefer  bool
}

//绑定
func (c *Container) Bind(abstract string, instance *BindNode)bool{
	c.locker.Lock()
	defer c.locker.Unlock()
	if c.bindList[abstract] == nil{
		c.bindList[abstract] = instance
		return true;
	}else{
		return false;
	}
}

//获取节点
func (c *Container) Get(abstract string) (interface{}, error) {

}
