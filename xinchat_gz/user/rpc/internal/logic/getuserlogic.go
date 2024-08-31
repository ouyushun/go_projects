package logic

import (
	"context"
	"fmt"

	"xinchat_gz/user/rpc/internal/svc"
	"xinchat_gz/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserRequest) (*user.GetUserResponse, error) {
	// todo: add your logic here and delete this line

	fmt.Println(in)
	res := &user.GetUserResponse{Id: in.Id, Name: "aaa", Phone: "9999999"}
	return res, nil
}
