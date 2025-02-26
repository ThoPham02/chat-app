// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6

package handler

import (
	"net/http"

	"chat_app/service/chat/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserTokenMiddleware},
			[]rest.Route{
				{
					// Get list messages
					Method:  http.MethodGet,
					Path:    "/:conversationID/messages",
					Handler: GetListMessageHandler(serverCtx),
				},
				{
					// Send message
					Method:  http.MethodPost,
					Path:    "/:conversationID/messages",
					Handler: SendMessageHandler(serverCtx),
				},
				{
					// Find messages
					Method:  http.MethodGet,
					Path:    "/:conversationID/messages-find",
					Handler: FindMessageHandler(serverCtx),
				},
				{
					// Create a new conversations
					Method:  http.MethodPost,
					Path:    "/conversations",
					Handler: CreateConversationHandler(serverCtx),
				},
				{
					// Get list conversations
					Method:  http.MethodGet,
					Path:    "/conversations",
					Handler: GetListConversationHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/chats"),
	)
}
