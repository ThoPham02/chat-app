package handler

import (
	"net/http"

	"chat_app/service/chat/api/internal/logic"
	"chat_app/service/chat/api/internal/svc"
	"chat_app/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetListConversationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetListConversationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetListConversationLogic(r.Context(), svcCtx)
		resp, err := l.GetListConversation(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
