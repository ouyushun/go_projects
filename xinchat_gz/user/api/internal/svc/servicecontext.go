package svc

import (
	"github.com/kr/pretty"
	"github.com/zeromicro/go-zero/zrpc"
	"xinchat_gz/user/api/internal/config"
	"xinchat_gz/user/rpc/userclient"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	pretty.Println(c)
	return &ServiceContext{
		Config:        c,
		UserRpcClient: userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
