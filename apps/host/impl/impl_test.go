package impl_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/renmcc/go-cmdb/apps/host"
	"github.com/renmcc/go-cmdb/conf"
	"github.com/stretchr/testify/assert"
)

var (
	// 定义对象是满足该接口的实现
	service host.Service
)

func TestCreate(t *testing.T) {
	should := assert.New(t)
	ins := host.NewHost()
	ins.Id = "ins-03"
	ins.Name = "test"
	ins.Region = "cn-杭州"
	ins.Type = "sm1"
	ins.CPU = 1
	ins.Memory = 2048
	// 查看结构体数据
	typeInfo := reflect.TypeOf(*ins)
	valInfo := reflect.ValueOf(*ins)
	for i := 0; i < typeInfo.NumField(); i++ {
		key := typeInfo.Field(i).Name
		value := valInfo.Field(i).Interface()
		fmt.Printf("%+v:%+v\n", key, value)
	}
	ins, err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}
}

func init() {
	// 初始化全局Logger
	zap.DevelopmentSetup()
	// 初始化config
	conf.LoadConfigFromToml("../../../etc/config.toml")
	// host service 的具体实现
	// service = impl.NewMysqlImpl()

}
