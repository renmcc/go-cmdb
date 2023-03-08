package host

import (
	"context"
)

// host app  增删改查接口定义
type Service interface {
	// 录入主机
	CreateHost(context.Context, *Host) error
	// 查询主机列表
	ListHost(context.Context, *ListHostRequest) (*HostSet, error)
	// 查询主机详情
	QueryHost(context.Context, *QueryHostRequest) (*Host, error)
	// 主机更新
	UpdateHost(context.Context, *Host) (*Host, error)
	// 主机删除
	DeleteHost(context.Context, *DeleteHostRequest) error
}

type HostSet struct {
	Total int     `json:"total"`
	Items []*Host `json:"items"`
}

func (s *HostSet) Add(item *Host) {
	s.Items = append(s.Items, item)
}

// 构造函数
func NewHostSet() *HostSet {
	return &HostSet{
		Items: []*Host{},
		Total: 0,
	}
}

type ListHostRequest struct {
	PageSize     int    `json:"page_size" validate:"max=50"`
	PageNumber   int    `json:"page_number" `
	SerialNumber string `json:"serial_number" validate:"max=15"`
	PrivateIp    string `json:"privateip" validate:"max=16"`
}

func (req *ListHostRequest) GetPageSize() uint {
	return uint(req.PageSize)
}

func (req *ListHostRequest) OffSet() int64 {
	return int64((req.PageNumber - 1) * req.PageSize)
}

// 结构体校验
func (req *ListHostRequest) Validate() error {
	return validate.Struct(req)
}

// 构造函数
func NewListHostRequest() *ListHostRequest {
	return &ListHostRequest{
		PageSize:     20,
		PageNumber:   1,
		SerialNumber: "%",
		PrivateIp:    "%",
	}
}

type QueryHostRequest struct {
	Xid string `json:"xid"  validate:"required"`
}

// 结构体校验
func (req *QueryHostRequest) Validate() error {
	return validate.Struct(req)
}

// 构造函数
func NewQueryHostRequest(xid string) *QueryHostRequest {
	return &QueryHostRequest{
		Xid: xid,
	}
}

type DeleteHostRequest struct {
	Xid string `json:"xid"  validate:"required"`
}

// 结构体校验
func (req *DeleteHostRequest) Validate() error {
	return validate.Struct(req)
}

// 构造函数
func NewDeleteHostRequest(xid string) *DeleteHostRequest {
	return &DeleteHostRequest{
		Xid: xid,
	}
}
