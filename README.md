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
│   ├── apps.go                               # app注册
│   ├── host                                  # 具体业务场景领域服务 
│   │   ├── app.go
│   │   ├── api                               # api接口
│   │   │   ├── http.go
│   │   │   └── view.go                       # api逻辑
│   │   ├── impl                             # 具体接口实现
│   │   │   ├── impl.go
│   │   │   ├── impl_test.go
│   │   │   ├── sql.go
│   │   │   └── view.go                      # curd逻辑
│   │   ├── interface.go                     # 标准
│   │   └── model.go                         # 数据模型
│   └── all
│       └── impl.go                          # 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载。    
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
│   ├── grpc.go                             
│   └── http.go
├── README.md
└── version                                  # 程序版本信息
    └── version.go
```

# 快速开发

1. 添加配置文件(默认读取位置: etc/config.toml)

2. 启动服务
```
# 编译protobuf文件, 生成代码
$ make dep
# 运行程序
$ make run
```