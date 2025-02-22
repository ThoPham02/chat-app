package logic

import (
	"context"

	"chat_app/service/chat/api/internal/svc"
	"chat_app/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListConversationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListConversationLogic {
	return &GetListConversationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListConversationLogic) GetListConversation(req *types.GetListConversationReq) (resp *types.GetListConversationRes, err error) {
	// todo: add your logic here and delete this line

	return
}
