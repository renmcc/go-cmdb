package model

type Base struct {
	Xid      string `json:"xid`        //唯一id
	Status   int    `json:"status"`    // 0删除 1创建 2更新
	CreateAt int64  `json:"create_at"` // 创建时间
	CreateBy string `json:"create_by"` // 创建操作人
	UpdateAt int64  `json:"update_at"` // 更新时间
	UpdateBy string `json:"update_by"` // 更新操作人
	DeleteAt int64  `json:"delete_at"` // 删除时间
	DeleteBy string `json:"delete_by"` // 删除操作人
}
