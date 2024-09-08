package logic

import (
	"context"
	"xinchat_gz/userdd/model"
	"xinchat_gz/userdd/rpc/internal/svc"
	"xinchat_gz/userdd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	// todo: add your logic here and delete this line

	name := in.Name
	password := in.Password
	phone := in.Phone

	newUser := &model.User{
		Name:     name,
		Password: password,
		Phone:    phone,
	}

	_, err := l.svcCtx.UserModel.Insert(l.ctx, newUser)

	if err != nil {
		return nil, err
	}
	return nil, nil
}
