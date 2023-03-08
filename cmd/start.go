package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/opengoats/goat/logger"
	"github.com/opengoats/goat/logger/zap"
	"github.com/renmcc/go-cmdb/apps"
	"github.com/renmcc/go-cmdb/protocol"

	"github.com/renmcc/go-cmdb/conf"
	"github.com/spf13/cobra"

	// 注册所有的app
	_ "github.com/renmcc/go-cmdb/apps/all"
)

var (
	// pusher service config option
	confFile string
)

// 用于管理所有需要启动的服务
type manager struct {
	http *protocol.HttpService
	log  logger.Logger
}

// 处理来自外部的中断信号, 比如Terminal
func (m *manager) WaitStop(ch <-chan os.Signal) {
	for v := range ch {
		switch v {
		default:
			m.log.Infof("received signal: %s", v)
			// 关闭外部调用
			m.http.Stop()
		}
	}
}

// 构造函数
func newManager() *manager {
	return &manager{
		http: protocol.NewHttpService(),
		log:  zap.L().Named("CLI"),
	}
}

// 程序的启动时 组装都在这里进行
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 cmdb-api",
	Long:  "启动 cmdb-api",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载程序配置
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			return err
		}
		// 初始化全局日志Logger
		if err := conf.LoadGlobalLogger(); err != nil {
			return err
		}
		// 所有注册服务进行初始化
		apps.InitImpl()

		// 启动http服务
		svc := newManager()
		ch := make(chan os.Signal, 1)
		defer close(ch)

		signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
		go svc.WaitStop(ch)
		return svc.http.Start()
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/config.toml", "配置文件")
	RootCmd.AddCommand(StartCmd)
}
