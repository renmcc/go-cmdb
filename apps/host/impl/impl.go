package impl

import (
	"database/sql"

	"github.com/renmcc/go-cmdb/apps"
	"github.com/renmcc/go-cmdb/apps/host"
	"github.com/renmcc/go-cmdb/conf"
	"github.com/renmcc/toolbox/logger"
	"github.com/renmcc/toolbox/logger/zap"
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

// 构造函数
func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		log: zap.L().Named("Host"),
		db:  conf.C().MySQL.GetDB(),
	}
}

// _ import app 自动执行注册逻辑
func init() {
	//  对象注册到ioc层
	apps.RegistryImpl(impl)
}
