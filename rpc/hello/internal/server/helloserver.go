// Code generated by goctl. DO NOT EDIT!
// Source: hello.proto

package server

import (
	"context"

	"github.com/1005281342/gozerodemo/rpc/hello/hello"
	"github.com/1005281342/gozerodemo/rpc/hello/internal/logic"
	"github.com/1005281342/gozerodemo/rpc/hello/internal/svc"
)

type HelloServer struct {
	svcCtx *svc.ServiceContext
}

func NewHelloServer(svcCtx *svc.ServiceContext) *HelloServer {
	return &HelloServer{
		svcCtx: svcCtx,
	}
}

func (s *HelloServer) Say(ctx context.Context, in *hello.SayReq) (*hello.SayRsp, error) {
	l := logic.NewSayLogic(ctx, s.svcCtx)
	return l.Say(in)
}
