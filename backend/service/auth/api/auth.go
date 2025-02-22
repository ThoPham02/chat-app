package api

import (
	"chat_app/service/auth/api/internal/config"
	"chat_app/service/auth/api/internal/handler"
	"chat_app/service/auth/api/internal/svc"
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("auth-api", "etc/auth-api.yaml", "the config file")

type AuthService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewAuthService(server *rest.Server) *AuthService {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	return &AuthService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (s *AuthService) Start() error {
	return nil
}
