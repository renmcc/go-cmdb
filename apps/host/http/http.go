package http

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/renmcc/go-cmdb/apps"
	"github.com/renmcc/go-cmdb/apps/host"
)

// 面向接口, 真正Service的实现, 在服务实例化的时候传递进行
// 也就是(CLI)  Start时候
var handler = &Handler{}

// 通过写一个实例类, 把内部的接口通过HTTP协议暴露出去
// 所以需要依赖内部接口的实现
// 该实体类, 会实现Gin的Http Handler
type Handler struct {
	svc host.Service
	log logger.Logger
}

// 完成了 Http Handler的注册
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
	r.GET("/hosts", h.queryHost)
	r.GET("/hosts/:id", h.describeHost)
	r.PUT("/hosts/:id", h.putHost)
}

func (h *Handler) Name() string {
	return host.AppName
}

// 从IOC里面获取HostService的实例对象,进行断言
func (h *Handler) Config() {
	h.svc = apps.GetImpl(h.Name()).(host.Service)
	h.log = zap.L().Named("Host.Http")
}

// 完成Http Handler注册
func init() {
	apps.RegistryGin(handler)
}
