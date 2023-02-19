package host

import "context"

// host app  增删改查接口定义
type Service interface {
	// 录入主机
	CreateHost(context.Context, *Host) (*Host, error)
	// 查询主机列表
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	// 查询主机详情
	DescribeHost(context.Context, *QueryHostRequest) (*Host, error)
	// 主机更新
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
	// 主机删除
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
}

type HostSet struct {
	Items []*Host
	Total int
}

type QueryHostRequest struct {
}

type UpdateHostRequest struct {
	*Describe
}

type DeleteHostRequest struct {
	Id string
}
