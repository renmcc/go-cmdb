package impl_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/opengoats/goat/logger/zap"
	"github.com/renmcc/go-cmdb/apps/host"
	"github.com/renmcc/go-cmdb/apps/host/impl"
	"github.com/renmcc/go-cmdb/conf"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

var (
	// 定义对象是满足该接口的实现
	service host.Service
)

func TestCreate(t *testing.T) {
	should := assert.New(t)
	ins := host.NewHost()
	ins.ResourceId = "xxxxxxxxxxxxx"
	ins.Vendor = 1
	ins.Name = "测试ecs"
	ins.Region = "cn-杭州"
	ins.PublicIP = "1.1.1.1"
	ins.PrivateIP = "2.2.2.2"
	ins.CPU = "4"
	ins.Memory = "1028"
	ins.OSType = "linux"
	ins.OSName = "centos7"
	ins.SerialNumber = xid.New().String()
	// 查看结构体数据
	typeInfo := reflect.TypeOf(*ins)
	valInfo := reflect.ValueOf(*ins)
	for i := 0; i < typeInfo.NumField(); i++ {
		key := typeInfo.Field(i).Name
		value := valInfo.Field(i).Interface()
		fmt.Printf("%+v:%+v\n", key, value)
	}
	err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(err)
	}
}

func TestListHost(t *testing.T) {
	should := assert.New(t)
	req := host.NewListHostRequest()
	set, err := service.ListHost(context.Background(), req)
	if should.NoError(err) {
		for _, v := range set.Items {
			fmt.Printf("v: %+v\n", *v)
		}
		fmt.Println(set.Total)
	}
}

func TestQueryHost(t *testing.T) {
	should := assert.New(t)

	req := host.NewQueryHostRequest("ins-05")
	ins, err := service.QueryHost(context.Background(), req)
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
	service = impl.NewHostServiceImpl()

}
