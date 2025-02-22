package api

import (
	"flag"

	"chat_app/service/chat/api/internal/config"
	"chat_app/service/chat/api/internal/handler"
	"chat_app/service/chat/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("chat-api", "etc/chat-api.yaml", "the config file")

type ChatService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewChatService(server *rest.Server) *ChatService {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	return &ChatService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (s *ChatService) Start() error {
	return nil
}
