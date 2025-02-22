package logic

import (
	"context"

	"chat_app/service/chat/api/internal/svc"
	"chat_app/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMessageLogic {
	return &FindMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindMessageLogic) FindMessage(req *types.GetFindMessageReq) (resp *types.GetFindMessageRes, err error) {
	// todo: add your logic here and delete this line

	return
}
