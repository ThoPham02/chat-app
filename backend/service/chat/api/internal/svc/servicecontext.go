package svc

import (
	"chat_app/service/chat/api/internal/config"
	"chat_app/service/chat/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config              config.Config
	UserTokenMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		UserTokenMiddleware: middleware.NewUserTokenMiddleware().Handle,
	}
}
