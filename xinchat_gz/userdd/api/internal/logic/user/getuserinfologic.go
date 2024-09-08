package user

import (
	"context"
	"xinchat_gz/userdd/rpc/userclient"

	"xinchat_gz/userdd/api/internal/svc"
	"xinchat_gz/userdd/api/internal/types"

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
	user, err := l.svcCtx.UserRpcClient.GetUser(l.ctx, &userclient.GetUserRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.GetUserInfoResp{
		Id:   user.Id,
		Name: user.Name,
	}, err
}
