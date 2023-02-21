package all

// 所有模块的注册
import (
	_ "github.com/renmcc/go-cmdb/apps/host/http"
	_ "github.com/renmcc/go-cmdb/apps/host/impl"
	// _ "github.com/renmcc/go-cmdb/apps/host/sqlitImpl"
)
