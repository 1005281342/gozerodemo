package handler

import (
	"net/http"

	"github.com/1005281342/gozerodemo/api/internal/logic"
	"github.com/1005281342/gozerodemo/api/internal/svc"
	"github.com/1005281342/gozerodemo/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func SayHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SayReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSayLogic(r.Context(), ctx)
		resp, err := l.Say(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
