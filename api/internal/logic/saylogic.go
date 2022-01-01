package logic

import (
	"context"
	"github.com/1005281342/gozerodemo/rpc/hello/hello"
	"net/http"

	"github.com/1005281342/gozerodemo/api/internal/svc"
	"github.com/1005281342/gozerodemo/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSayLogic(ctx context.Context, svcCtx *svc.ServiceContext) SayLogic {
	return SayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SayLogic) Say(req types.SayReq) (*types.Rsp, error) {
	var rsp, err = l.svcCtx.HelloRPC.Say(l.ctx, &hello.SayReq{Name: req.Name})
	if err != nil {
		return nil, err
	}

	return &types.Rsp{Data: rsp.GetReply(), Result: true, Code: http.StatusOK}, nil
}
