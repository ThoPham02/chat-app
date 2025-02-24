package svc

import (
	"chat_app/service/chat/api/internal/config"
	"chat_app/service/chat/api/internal/middleware"
	"chat_app/service/chat/model"
	"chat_app/sync"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config                  config.Config
	UserTokenMiddleware     rest.Middleware
	ObjSync                 sync.ObjSync
	ConversationModel       model.ConversationsModel
	ConversationMemberModel model.ConversationMembersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                  c,
		UserTokenMiddleware:     middleware.NewUserTokenMiddleware().Handle,
		ObjSync:                 *sync.NewObjSync(0),
		ConversationModel:       model.NewConversationsModel(sqlx.NewMysql(c.DataSource)),
		ConversationMemberModel: model.NewConversationMembersModel(sqlx.NewMysql(c.DataSource)),
	}
}
