package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"xinchat_gz/userdd/model"
	"xinchat_gz/userdd/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
