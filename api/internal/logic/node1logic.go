package logic

import (
	"context"
	"github.com/1005281342/gozerodemo/rpc/jaeger/jaeger"
	"net/http"

	"github.com/1005281342/gozerodemo/api/internal/svc"
	"github.com/1005281342/gozerodemo/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type Node1Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNode1Logic(ctx context.Context, svcCtx *svc.ServiceContext) Node1Logic {
	return Node1Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Node1Logic) Node1(req types.Req) (*types.Rsp, error) {
	var s = NewSayLogic(l.ctx, l.svcCtx)
	s.Say(types.SayReq{Name: req.Message})

	var rsp, err = l.svcCtx.JaegerRPC.Node1(l.ctx, &jaeger.NodeReq{Message: req.Message, Cnt: req.Cnt})
	if err != nil {
		return nil, err
	}

	return &types.Rsp{Code: http.StatusOK, Result: true, Data: rsp.Message}, nil
}
