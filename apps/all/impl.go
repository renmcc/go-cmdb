package all

// 所有模块的注册
import (
	_ "github.com/renmcc/go-cmdb/apps/book/api"
	_ "github.com/renmcc/go-cmdb/apps/book/impl"
	_ "github.com/renmcc/go-cmdb/apps/host/api"
	_ "github.com/renmcc/go-cmdb/apps/host/impl"
)
