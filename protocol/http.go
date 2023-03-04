package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opengoats/goat/logger"
	"github.com/opengoats/goat/logger/zap"
	"github.com/renmcc/go-cmdb/apps"
	"github.com/renmcc/go-cmdb/conf"
)

type HttpService struct {
	server *http.Server
	log    logger.Logger
	router gin.IRouter
}

func (s *HttpService) Start() error {
	// 加载Handler, 把所有的模块的Handler注册给了Gin Router
	apps.InitGin(s.router)

	// 已加载App的日志信息
	apps := apps.LoadedGinApps()
	s.log.Infof("loaded gin apps :%v", apps)

	// 该操作时阻塞的, 简单端口，等待请求
	// 如果服务的正常关闭,
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.log.Info("service stoped success")
			return nil
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}

	return nil
}

func (s *HttpService) Stop() {
	s.log.Info("start graceful shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.log.Warnf("shut down http service error, %s", err)
	}
}

// HttpService构造函数
func NewHttpService() *HttpService {
	// 关闭debug
	// gin.SetMode(gin.ReleaseMode)
	// new gin router实例, 并没有加载Handler
	r := gin.Default()
	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.HttpAddr(),
		Handler:           r,
	}
	return &HttpService{
		server: server,
		log:    zap.L().Named("HTTP Service"),
		router: r,
	}
}
