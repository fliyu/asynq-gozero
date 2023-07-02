package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero-asynq/asynqclient/internal/logic"
	"gozero-asynq/asynqclient/internal/svc"
	"gozero-asynq/asynqclient/internal/types"
)

func AsynqclientHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAsynqclientLogic(r.Context(), svcCtx)
		resp, err := l.Asynqclient(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
