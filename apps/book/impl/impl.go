package impl

import (
	"database/sql"

	"github.com/opengoats/goat/logger"
	"github.com/opengoats/goat/logger/zap"
	"github.com/renmcc/go-cmdb/apps"
	"github.com/renmcc/go-cmdb/apps/book"
	"github.com/renmcc/go-cmdb/conf"
	"google.golang.org/grpc"
)

var (
	// Service 服务实例
	svr                    = &service{}
	_   book.ServiceServer = svr
)

// GRPC接口的实现类
type service struct {
	db   *sql.DB
	log  logger.Logger
	book book.ServiceServer
	book.UnimplementedServiceServer
}

func (s *service) Config() {
	s.log = zap.L().Named(s.Name())
	s.db = conf.C().MySQL.GetDB()
	s.book = apps.GetGrpcImpl(book.AppName).(book.ServiceServer)
}

func (s *service) Name() string {
	return book.AppName
}

func (s *service) Registry(server *grpc.Server) {
	book.RegisterServiceServer(server, svr)
}

func init() {
	apps.RegistryGrpc(svr)
}
