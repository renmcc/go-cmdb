package host

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/renmcc/go-cmdb/common/model"
	"github.com/rs/xid"
)

var (
	validate = validator.New()
)

type Vendor int

const (
	// 枚举的默认值
	PRIVATE_IDC Vendor = iota
	// 阿里云
	ALIYUN
	// 腾讯云
	TXYUN
)

// Host模型的定义
type Host struct {
	*model.Base
	ResourceId   string `json:"resource_id"  validate:"required,max=12"` // 全局唯一Id
	Vendor       Vendor `json:"vendor"`                                  // 厂商
	Name         string `json:"name"  validate:"required,max=12" `       // 名称
	Region       string `json:"region"  validate:"required,max=12"`      // 地域
	ExpireAt     int64  `json:"expire_at"`                               // 过期时间
	Description  string `json:"description"`                             // 描述
	PublicIP     string `json:"public_ip"`                               // 公网IP
	PrivateIP    string `json:"private_ip" validate:"required"`          // 内网IP
	CPU          string `json:"cpu"`                                     // 核数
	Memory       string `json:"memory"`                                  // 内存
	OSType       string `json:"os_type"`                                 // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`                                 // 操作系统名称
	SerialNumber string `json:"serial_number" validate:"required"`       // 序列号
}

// 结构体校验
func (h *Host) Validate() error {
	return validate.Struct(h)
}

// 插入默认值
func (h *Host) InsertDefault() {
	h.Base.Xid = xid.New().String()
	h.Base.CreateAt = time.Now().UnixMicro()
	h.Base.CreateBy = ""
	h.Base.Status = 1
}

// 更新默认值
func (h *Host) UpdateDefault() {
	h.Base.UpdateAt = time.Now().UnixMicro()
	h.Base.UpdateBy = ""
}

// 删除默认值
func (h *Host) DeleteDefault() {
	h.Base.DeleteAt = time.Now().UnixMicro()
	h.Base.DeleteBy = ""
}

// 对象全量更新
func (h *Host) Put(obj *Host) error {
	if obj.Xid != h.Xid {
		return fmt.Errorf("id not equal")
	}
	*h = *obj
	return nil
}

// Host模型构造函数
func NewHost() *Host {
	return &Host{
		Base: &model.Base{},
	}
}
