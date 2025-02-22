package logic

import (
	"context"

	"chat_app/service/chat/api/internal/svc"
	"chat_app/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateConversationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateConversationLogic {
	return &CreateConversationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateConversationLogic) CreateConversation(req *types.CreateConversationReq) (resp *types.CreateConversationRes, err error) {
	// todo: add your logic here and delete this line

	return
}
