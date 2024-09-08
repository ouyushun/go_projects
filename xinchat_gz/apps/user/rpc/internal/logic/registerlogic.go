package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"time"
	"xinchat_gz/pkg/ctxdata"
	"xinchat_gz/pkg/encrypt"
	"xinchat_gz/pkg/wuid"

	"xinchat_gz/apps/user/models"
	"xinchat_gz/apps/user/rpc/internal/svc"
	"xinchat_gz/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrorPhoneExist = errors.New("手机号码已注册")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line
	UserEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)

	//查到记录
	if UserEntity != nil {
		return nil, ErrorPhoneExist
	}
	if err != nil && err != models.ErrNotFound {
		return nil, err
	}

	newUser := &models.Users{
		Id:       wuid.GenUid(l.svcCtx.Config.Mysql.DataSource),
		Avatar:   in.Avatar,
		Nickname: in.Nickname,
		Phone:    in.Phone,
		Sex: sql.NullInt64{
			Int64: int64(in.Sex),
			Valid: true,
		},
	}

	if len(in.Password) > 0 {
		encryptPassword, err := encrypt.GenPasswordHash([]byte(in.Password))
		if err != nil {
			return nil, err
		}
		newUser.Password = sql.NullString{
			String: string(encryptPassword),
			Valid:  true,
		}
	}

	_, err = l.svcCtx.UsersModel.Insert(l.ctx, newUser)
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, l.svcCtx.Config.Jwt.AccessExpire, newUser.Id)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
