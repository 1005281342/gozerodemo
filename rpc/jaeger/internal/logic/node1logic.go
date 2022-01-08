package logic

import (
	"context"
	"github.com/1005281342/gozerodemo/rpc/hello/helloclient"

	"github.com/1005281342/gozerodemo/rpc/jaeger/internal/svc"
	"github.com/1005281342/gozerodemo/rpc/jaeger/jaeger"

	"github.com/tal-tech/go-zero/core/logx"
)

type Node1Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNode1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Node1Logic {
	return &Node1Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Node1Logic) Node1(in *jaeger.NodeReq) (*jaeger.NodeRsp, error) {

	in.Cnt--
	in.Message = in.GetMessage() + "node1"
	if in.Cnt == 0 {
		return &jaeger.NodeRsp{Message: in.Message}, nil
	}
	if in.Cnt&1 == 0 {
		var rsp, err = l.svcCtx.HelloRPC.Say(l.ctx, &helloclient.SayReq{Name: in.Message})
		if err != nil {
			in.Message += err.Error()
		} else {
			in.Message += rsp.GetReply()
		}
		var n = NewNode2Logic(l.ctx, l.svcCtx)
		return n.Node2(in)
	}

	var n = NewNode3Logic(l.ctx, l.svcCtx)
	return n.Node3(in)
}
