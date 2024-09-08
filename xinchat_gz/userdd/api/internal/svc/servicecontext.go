package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"xinchat_gz/userdd/api/internal/config"
	"xinchat_gz/userdd/api/internal/middleware"
	"xinchat_gz/userdd/rpc/userclient"
)

type ServiceContext struct {
	Config            config.Config
	UserRpcClient     userclient.User
	Loginverification rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		UserRpcClient:     userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		Loginverification: middleware.NewLoginverificationMiddleware().Handle,
	}
}
