package logic

import (
	"context"
	"github.com/1005281342/gozerodemo/rpc/hello/hello"

	"github.com/1005281342/gozerodemo/rpc/hello/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type SayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayLogic {
	return &SayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayLogic) Say(in *hello.SayReq) (*hello.SayRsp, error) {
	return &hello.SayRsp{Reply: "hello, " + in.Name}, nil
}
