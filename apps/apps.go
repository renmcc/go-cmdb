package apps

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// IOC 容器层: 管理所有app实现和gin路由

var (
	// 维护当前所有的服务
	implApps    = map[string]ImplService{}
	ginApps     = map[string]GinService{}
	restfulApps = map[string]RestfulService{}
	grpcApps    = map[string]GrpcService{}
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

type GrpcService interface {
	Registry(g *grpc.Server)
	Config()
	Name() string
}

func RegistryGrpc(svc GrpcService) {
	// 服务实例注册到svcs map当中
	if _, ok := grpcApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	grpcApps[svc.Name()] = svc
}

// 注册到Ioc容器里面的所有gin路由
func InitGrpc(r *grpc.Server) {
	// 先初始化好所有对象
	for _, v := range grpcApps {
		v.Config()
	}

	// 完成Grpc的注册
	for _, v := range grpcApps {
		v.Registry(r)
	}
}

// 已经加载完成的Grpc App由哪些
func LoadedGrpcApps() (names []string) {
	for k := range grpcApps {
		names = append(names, k)
	}

	return
}

// 返回一个对象, 任何类型都可以, 使用时, 由使用方进行断言
func GetGrpcImpl(name string) interface{} {
	for k, v := range grpcApps {
		if k == name {
			return v
		}
	}

	return nil
}

type RestfulService interface {
	Registry(ws *restful.WebService)
	Config()
	Name() string
}

func RegistryRestful(svc RestfulService) {
	// 服务实例注册到svcs map当中
	if _, ok := restfulApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	restfulApps[svc.Name()] = svc
}

// 注册restful web service
// root router
func InitRestful(r *restful.Container) {
	// 先初始化好所有对象
	for _, v := range restfulApps {
		v.Config()
	}

	// 完成Grpc的注册
	for _, v := range restfulApps {
		ws := new(restful.WebService)
		r.Add(ws)
		v.Registry(ws)
	}
}

// 已经加载完成的Grpc App由哪些
func LoadedRestfulApps() (names []string) {
	for k := range restfulApps {
		names = append(names, k)
	}

	return
}
