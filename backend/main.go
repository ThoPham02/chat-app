package main

import (
	"chat_app/config"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"

	authApi "chat_app/service/auth/api"
	chatApi "chat_app/service/chat/api"
)

var configFile = flag.String("f", "etc/server.yaml", "the config file")

// @BasePath /
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))

	logx.DisableStat()
	defer server.Stop()

	authService := authApi.NewAuthService(server)
	authService.Start()
	authFunc := authApi.NewAuthFunction(authService)
	authFunc.Start()

	chatService := chatApi.NewChatService(server)
	chatService.Start()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
