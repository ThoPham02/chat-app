package logic

import (
	"context"

	"chat_app/common"
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
	l.Logger.Info("CreateConversation req: ", req)

	var conversation types.Conversation

	l.Logger.Info("CreateConversation Success")
	return &types.CreateConversationRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Conversation: conversation,
	}, nil
}
