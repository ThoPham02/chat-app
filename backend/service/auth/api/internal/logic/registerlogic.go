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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	l.Logger.Info("Register req: ", req)

	var accessExpire = l.svcCtx.Config.Auth.AccessExpire
	var accessSecret = l.svcCtx.Config.Auth.AccessSecret
	var now int64 = time.Now().Unix()

	var userID int64 = l.svcCtx.ObjSync.GenServiceObjID()
	var hashedPassword string
	var currentTime = common.GetCurrentTime()
	var userModel *model.Users

	userModel, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.UserName)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if userModel != nil {
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.USER_EXISTED_CODE,
				Message: common.USER_EXISTED_MESS,
			},
		}, nil
	}

	hashedPassword, err = util.HashPassword(req.Password)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.UNKNOW_CODE,
				Message: common.UNKNOW_MESS,
			},
		}, nil
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.Users{
		UserId:    userID,
		Username:  req.UserName,
		HashPw:    hashedPassword,
		Fullname:  req.Fullname,
		AvatarUrl: sql.NullString{Valid: true, String: ""},
		CreatedAt: sql.NullInt64{Valid: true, Int64: currentTime},
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	token, err := util.GenerateToken(accessSecret, now, accessExpire, userID)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.UNKNOW_CODE,
				Message: common.UNKNOW_MESS,
			},
		}, nil
	}

	l.Logger.Info("Register success")
	return &types.RegisterRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Token: token,
		User: types.User{
			UserID:    userID,
			FullName:  req.Fullname,
			AvatarUrl: "",
			CreatedAt: currentTime,
		},
	}, nil
}
