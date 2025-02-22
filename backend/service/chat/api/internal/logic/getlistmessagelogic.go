package logic

import (
	"context"

	"chat_app/service/chat/api/internal/svc"
	"chat_app/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListMessageLogic {
	return &GetListMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListMessageLogic) GetListMessage(req *types.GetListMessageReq) (resp *types.GetListMessageRes, err error) {
	// todo: add your logic here and delete this line

	return
}
