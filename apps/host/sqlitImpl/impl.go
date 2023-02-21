package sqlitImpl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/renmcc/go-cmdb/apps"
	"github.com/renmcc/go-cmdb/apps/host"
)

var (
	_    host.Service = &HostServiceImpl{}
	impl              = &HostServiceImpl{}
)

type HostServiceImpl struct {
	log logger.Logger
}

// 服务的名称
func (i *HostServiceImpl) Name() string {
	return host.AppName
}

// 调用需保证全局对象Config和全局Logger已经初始化完成
func (i *HostServiceImpl) Config() {
	i.log = zap.L().Named("Host.Sqllit")
}

// _ import app 自动执行注册逻辑
func init() {
	//  对象注册到ioc层
	apps.RegistryImpl(impl)
}
