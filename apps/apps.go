package apps

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// IOC 容器层: 管理所有app实现和gin路由

var (
	// 维护当前所有的服务
	implApps = map[string]ImplService{}
	ginApps  = map[string]GinService{}
)

// 注册所有的app
type ImplService interface {
	Config()
	Name() string
}

func RegistryImpl(svc ImplService) {
	// 判断是否已经注册
	if _, ok := implApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}
	// 服务实例注册到svcs map当中
	implApps[svc.Name()] = svc
}

// app初始化
func InitImpl() {
	for _, v := range implApps {
		v.Config()
	}
}

// 返回一个对象, 任何类型都可以, 使用时, 由使用方进行断言
func GetImpl(name string) interface{} {
	for k, v := range implApps {
		if k == name {
			return v
		}
	}

	return nil
}

// 注册Gin编写的Handler
type GinService interface {
	Registry(r gin.IRouter)
	Config()
	Name() string
}

func RegistryGin(svc GinService) {
	// 服务实例注册到svcs map当中
	if _, ok := ginApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	ginApps[svc.Name()] = svc
}

// 注册到Ioc容器里面的所有gin路由
func InitGin(r gin.IRouter) {
	// 先初始化好所有对象
	for _, v := range ginApps {
		v.Config()
	}

	// 完成Http Handler的注册
	for _, v := range ginApps {
		v.Registry(r)
	}
}

// 已经加载完成的Gin App由哪些
func LoadedGinApps() (names []string) {
	for k := range ginApps {
		names = append(names, k)
	}

	return
}
