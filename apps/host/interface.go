package host

import (
	"context"
)

// host app  增删改查接口定义
type Service interface {
	// 录入主机
	CreateHost(context.Context, *Host) (*Host, error)
	// 查询主机列表
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	// 查询主机详情
	DescribeHost(context.Context, *DescribeHostRequest) (*Host, error)
	// 主机更新
	UpdateHost(context.Context, *Host) (*Host, error)
	// 主机删除
	DeleteHost(context.Context, *DeleteHostRequest) error
}

type HostSet struct {
	Total int
	Items []*Host
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

type QueryHostRequest struct {
	PageSize    int    `json:"page_size" validate:"max=50"`
	PageNumber  int    `json:"page_number" `
	Name        string `json:"name" validate:"max=12"`
	Description string `json:"description" validate:"max=12"`
	PrivateIp   string `json:"privateip" validate:"max=16"`
	PublicIp    string `json:"publicip" validate:"max=16"`
}

func (req *QueryHostRequest) GetPageSize() uint {
	return uint(req.PageSize)
}

func (req *QueryHostRequest) OffSet() int64 {
	return int64((req.PageNumber - 1) * req.PageSize)
}

// 结构体校验
func (req *QueryHostRequest) Validate() error {
	return validate.Struct(req)
}

// 构造函数
func NewQueryHostRequest() *QueryHostRequest {
	return &QueryHostRequest{
		PageSize:    20,
		PageNumber:  1,
		Name:        "%",
		Description: "%",
		PrivateIp:   "%",
		PublicIp:    "%",
	}
}

type DescribeHostRequest struct {
	Id string `json:"id"  validate:"required,max=15"`
}

// 结构体校验
func (req *DescribeHostRequest) Validate() error {
	return validate.Struct(req)
}

// 构造函数
func NewDescribeHostRequest(id string) *DescribeHostRequest {
	return &DescribeHostRequest{
		Id: id,
	}
}

type DeleteHostRequest struct {
	Id string `json:"id"  validate:"required,max=15"`
}

// 结构体校验
func (req *DeleteHostRequest) Validate() error {
	return validate.Struct(req)
}

// 构造函数
func NewDeleteHostRequest(id string) *DeleteHostRequest {
	return &DeleteHostRequest{
		Id: id,
	}
}
