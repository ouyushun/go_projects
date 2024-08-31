package user

import (
	"context"

	"xinchat_gz/demo/internal/svc"
	"xinchat_gz/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserLogic {
	return &DelUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelUserLogic) DelUser(req *types.UpdateUserInfoReq) error {
	// todo: add your logic here and delete this line

	return nil
}
