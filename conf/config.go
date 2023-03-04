package conf

import (
	"database/sql"
	"fmt"
	"sync"
)

// 全局config实例对象,
// 也就是我们程序，在内存中的配置对象
// 程序内部获取配置, 都通过读取该对象
// 该Config对象 什么时候被初始化喃?
//
//	配置加载时:
//	   LoadConfigFromToml
//	   LoadConfigFromEnv
//
// 为了不被程序在运行时恶意修改, 设置成私有变量
var config *Config

// 全局MySQL 客户端实例
var db *sql.DB

// 要想获取配置, 单独提供函数
// 全局Config对象获取函数
func C() *Config {
	return config
}

// Config 应用配置
// 通过封装为一个对象, 来与外部配置进行对接
type Config struct {
	App   *App   `toml:"app"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}

// 初始化一个有默认值的Config对象
func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		Log:   NewDefaultLog(),
		MySQL: NewDefaultMySQL(),
	}
}

type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
}

// gin启动参数
func (a *App) HttpAddr() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

func (a *App) RestAddr() string {
	return fmt.Sprintf("%s:%s", a.Host, fmt.Sprintf("2%s", a.Port))
}

// grpc启动参数
func (a *App) GrpcAddr() string {
	return fmt.Sprintf("%s:%s", a.Host, fmt.Sprintf("1%s", a.Port))
}

// 构造函数
func NewDefaultApp() *App {
	return &App{
		Name: "demo",
		Host: "127.0.0.1",
		Port: "8050",
	}
}

// 用于配置全局Logger对象
type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
}

// 构造函数
func NewDefaultLog() *Log {
	return &Log{
		Level:  "info",
		Format: TextFormat,
		To:     ToStdout,
	}
}

// 用于配置全局MySql对象
type MySQL struct {
	Host        string `toml:"host" env:"D_MYSQL_HOST"`
	Port        string `toml:"port" env:"D_MYSQL_PORT"`
	UserName    string `toml:"username" env:"D_MYSQL_USERNAME"`
	Password    string `toml:"password" env:"D_MYSQL_PASSWORD"`
	Database    string `toml:"database" env:"D_MYSQL_DATABASE"`
	MaxOpenConn int    `toml:"max_open_conn" env:"D_MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int    `toml:"max_idle_conn" env:"D_MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int    `toml:"max_life_time" env:"D_MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int    `toml:"max_idle_time" env:"D_MYSQL_MAX_idle_TIME"`
	lock        sync.Mutex
}

// 构造函数
func NewDefaultMySQL() *MySQL {
	return &MySQL{
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "demo",
		Password:    "123456",
		Database:    "demo",
		MaxOpenConn: 200,
		MaxIdleConn: 100,
	}
}
