package impl

import (
	"database/sql"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/renmcc/go-cmdb/apps"
	"github.com/renmcc/go-cmdb/apps/host"
	"github.com/renmcc/go-cmdb/conf"
)

var (
	// 用于注册
	impl = &HostServiceImpl{}
	//确保接口实现
	_ host.Service = impl
)

// mysql实现类
type HostServiceImpl struct {
	db  *sql.DB
	log logger.Logger
}

// 服务的名称
func (i *HostServiceImpl) Name() string {
	return host.AppName
}

// 调用需保证全局对象Config和全局Logger已经初始化完成
func (i *HostServiceImpl) Config() {
	i.db = conf.C().MySQL.GetDB()
	i.log = zap.L().Named("Host")
}

// _ import app 自动执行注册逻辑
func init() {
	//  对象注册到ioc层
	apps.RegistryImpl(impl)
}