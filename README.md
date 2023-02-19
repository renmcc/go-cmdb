# go-cmdb
多云管理(完善中...)

# 如何启动
* 先启动注册中心
* 再启动用户中心
* 再启动CMDB

# 架构图
```
go-cmdb 
├── apps                                      # 具体业务场景的领域包
│   ├── apps.go
│   ├── host                                  # 具体业务场景领域服务 
│   │   ├── app.go
│   │   ├── http
│   │   │   ├── http.go
│   │   │   └── view.go
│   │   ├── impl                             # 做了一个mysql存储实现
│   │   │   ├── impl.go
│   │   │   ├── impl_test.go
│   │   │   ├── sql.go
│   │   │   └── view.go
│   │   ├── interface.go                     # host接口
│   │   └── model.go                         # host主机数据模型
│   └── registry
│       └── registry.go                      # 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载。    
├── cmd                                      # 脚手架功能: 处理程序启停参数，加载系统配置文件
│   ├── root.go
│   └── start.go
├── conf                                     # 脚手架功能: 配置文件加载
│   ├── config.go
│   ├── config_test.go
│   ├── db.go
│   ├── load.go
│   └── log.go
├── etc                                      # 配置文件
│   ├── config.toml
│   └── unit-test.env
├── go.mod                                   # go mod 依赖定义
├── go.sum
├── LICENSE
├── logs
│   └── api.log
├── main.go
├── Makefile                                 # make 命令定义
├── protocol                                 # 脚手架功能: rpc / http 功能加载
│   ├── grpc.go                              # 暂未实现
│   └── http.go
├── README.md
└── version                                  # 程序版本信息
    └── version.go
```

# 快速开发

make脚手架

$ make help
```
dep                            Get the dependencies
lint                           Lint Golang files
vet                            Run go vet
test                           Run unittests
test-coverage                  Run tests with coverage
build                          Build the binary file
linux                          Build the binary file
clean                          Remove previous build
help                           Display this help screen
```

1. 添加配置文件(默认读取位置: etc/config.toml)

2. 启动服务
```
# 编译protobuf文件, 生成代码
$ make gen
# 如果是MySQL, 执行SQL语句(docs/schema/tables.sql)
$ make init (未完成)
# 下载项目的依赖
$ make dep
# 运行程序
$ make run
```