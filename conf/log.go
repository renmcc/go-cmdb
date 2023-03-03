package conf

import (
	"fmt"

	"github.com/renmcc/toolbox/logger/zap"
)

// LogFormat 日志格式
type LogFormat string

const (
	// TextFormat 文本格式
	TextFormat = LogFormat("text")
	// JSONFormat json格式
	JSONFormat = LogFormat("json")
)

// LogTo 日志记录到哪儿
type LogTo string

const (
	// ToFile 保存到文件
	ToFile = LogTo("file")
	// ToStdout 打印到标准输出
	ToStdout = LogTo("stdout")
)

// log 为全局变量, 只需要load 即可全局可用户, 依赖全局配置先初始化
func LoadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)

	// 根据Config里面的日志配置，来配置全局Logger对象
	lc := C().Log

	// 解析日志Level配置
	// DebugLevel: "debug",
	// InfoLevel:  "info",
	// WarnLevel:  "warning",
	// ErrorLevel: "error",
	// FatalLevel: "fatal",
	// PanicLevel: "panic",
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}

	// 使用默认配置初始化Logger的全局配置
	zapConfig := zap.DefaultConfig()

	// 配置日志的Level基本
	zapConfig.Level = level

	// 程序每启动一次, 不必都生成一个新日志文件
	zapConfig.Files.RotateOnStartup = true

	// 配置日志的输出方式
	switch lc.To {
	case ToStdout:
		// 把日志打印到标准输出
		zapConfig.ToStderr = true
		// 把日志输入输出到文件
		zapConfig.ToFiles = false
	case ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}

	// 配置日志的输出格式:
	switch lc.Format {
	case JSONFormat:
		zapConfig.JSON = true
	case TextFormat:
		zapConfig.JSON = false
	}

	// 把配置应用到全局Logger
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}

	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}
