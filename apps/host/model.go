package host

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
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
	*Resource
	*Describe
}

// 结构体校验
func (h *Host) Validate() error {
	return validate.Struct(h)
}

// 字段默认值
func (h *Host) InjectDefault() {
	if h.CreateAt == 0 {
		h.CreateAt = time.Now().UnixMilli()
	}
}

// 对象全量更新
func (h *Host) Put(obj *Host) error {
	if obj.Id != h.Id {
		return fmt.Errorf("id not equal")
	}

	*h.Resource = *obj.Resource
	*h.Describe = *obj.Describe
	return nil
}

// Host模型构造函数
func NewHost() *Host {
	return &Host{
		Resource: &Resource{},
		Describe: &Describe{},
	}
}

// 资源公共属性部分
type Resource struct {
	Id          string            `json:"id"  validate:"required,max=12"`     // 全局唯一Id
	Vendor      Vendor            `json:"vendor"`                             // 厂商
	Region      string            `json:"region"  validate:"required,max=12"` // 地域
	CreateAt    int64             `json:"create_at"`                          // 创建时间
	ExpireAt    int64             `json:"expire_at"`                          // 过期时间
	Type        string            `json:"type"  validate:"required,max=12" `  // 规格
	Name        string            `json:"name"  validate:"required,max=12" `  // 名称
	Description string            `json:"description"`                        // 描述
	Status      string            `json:"status"`                             // 服务商中的状态
	Tags        map[string]string `json:"tags"`                               // 标签
	UpdateAt    int64             `json:"update_at"`                          // 更新时间
	SyncAt      int64             `json:"sync_at"`                            // 同步时间
	Account     string            `json:"accout"`                             // 资源的所属账号
	PublicIP    string            `json:"public_ip"`                          // 公网IP
	PrivateIP   string            `json:"private_ip"`                         // 内网IP
}

// 资源独有属性部分
type Describe struct {
	CPU          int    `json:"cpu" validate:"required"`    // 核数
	Memory       int    `json:"memory" validate:"required"` // 内存
	GPUAmount    int    `json:"gpu_amount"`                 // GPU数量
	GPUSpec      string `json:"gpu_spec"`                   // GPU类型
	OSType       string `json:"os_type"`                    // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`                    // 操作系统名称
	SerialNumber string `json:"serial_number"`              // 序列号
}
