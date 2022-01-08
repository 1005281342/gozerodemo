package logic

import (
	"context"

	"github.com/1005281342/gozerodemo/rpc/jaeger/internal/svc"
	"github.com/1005281342/gozerodemo/rpc/jaeger/jaeger"

	"github.com/tal-tech/go-zero/core/logx"
)

type Node3Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNode3Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Node3Logic {
	return &Node3Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Node3Logic) Node3(in *jaeger.NodeReq) (*jaeger.NodeRsp, error) {
	in.Cnt--
	in.Message = in.GetMessage() + "node3"
	if in.Cnt == 0 {
		return &jaeger.NodeRsp{Message: in.Message}, nil
	}
	if in.Cnt&1 == 0 {
		var n = NewNode1Logic(l.ctx, l.svcCtx)
		return n.Node1(in)
	}

	var n = NewNode2Logic(l.ctx, l.svcCtx)
	return n.Node2(in)
}
