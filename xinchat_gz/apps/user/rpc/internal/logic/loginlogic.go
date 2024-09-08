package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"xinchat_gz/apps/user/models"
	"xinchat_gz/apps/user/rpc/user"
	"xinchat_gz/pkg/ctxdata"
	"xinchat_gz/pkg/encrypt"

	"xinchat_gz/apps/user/rpc/internal/svc"
)

var ErrPhoneNotRegister = errors.New("手机号码未注册")
var ErrPasswordError = errors.New("密码错误")

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// todo: add your logic here and delete this line

	UserEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return nil, ErrPhoneNotRegister
		}
		return nil, err
	}

	//密码验证
	if !encrypt.ValidatePasswordHash(in.Password, UserEntity.Password.String) {
		return nil, ErrPasswordError
	}

	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, l.svcCtx.Config.Jwt.AccessExpire, UserEntity.Id)
	if err != nil {
		return nil, err
	}

	return &user.LoginResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
