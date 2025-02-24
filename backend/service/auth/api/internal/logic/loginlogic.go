package logic

import (
	"context"
	"database/sql"
	"time"

	"chat_app/common"
	"chat_app/service/auth/api/internal/svc"
	"chat_app/service/auth/api/internal/types"
	"chat_app/service/auth/model"
	"chat_app/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	l.Logger.Info("Login request: ", req)

	var accessExpire = l.svcCtx.Config.Auth.AccessExpire
	var accessSecret = l.svcCtx.Config.Auth.AccessSecret
	var now int64 = time.Now().Unix()

	var userModel *model.Users

	userModel, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.UserName)
	if err != nil && err != sql.ErrNoRows {
		l.Logger.Error(err)
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if userModel == nil {
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.USER_NOT_FOUND_CODE,
				Message: common.USER_NOT_FOUND_MESS,
			},
		}, nil
	}

	if !util.VerifyPassword(req.Password, userModel.HashPw) {
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.PASSWORD_INCORRECT_CODE,
				Message: common.PASSWORD_INCORRECT_MESS,
			},
		}, nil
	}

	token, err := util.GenerateToken(accessSecret, now, accessExpire, userModel.UserId)
	if err != nil {
		l.Logger.Error(err)
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.UNKNOW_CODE,
				Message: common.UNKNOW_MESS,
			},
		}, nil
	}

	l.Logger.Info("Login Success")
	return &types.LoginRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Token: token,
		User: types.User{
			UserID:    userModel.UserId,
			FullName:  userModel.Fullname,
			AvatarUrl: userModel.AvatarUrl.String,
			CreatedAt: userModel.CreatedAt.Int64,
		},
	}, nil
}
