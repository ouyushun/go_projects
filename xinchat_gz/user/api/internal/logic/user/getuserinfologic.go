package user

import (
	"context"
	"fmt"
	"xinchat_gz/user/rpc/userclient"

	"xinchat_gz/user/api/internal/svc"
	"xinchat_gz/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	// todo: add your logic here and delete this line

	fmt.Println("aaaaaaaaaaaaaaaaaa")

	user, err := l.svcCtx.UserRpcClient.GetUser(l.ctx, &userclient.GetUserRequest{Id: req.Id})
	return &types.GetUserInfoResp{
		Id:   user.Id,
		Name: user.Name,
	}, err
}
