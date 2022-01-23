package logic

import (
	"context"
	"github.com/1005281342/gozerodemo/api/internal/metrics"
	"github.com/1005281342/gozerodemo/rpc/hello/hello"
	rkmid "github.com/rookie-ninja/rk-entry/middleware"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
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
	// go-zero默认读global相关数据，因此在这里更新rk的trace到全局
	otel.SetTracerProvider(l.ctx.Value(rkmid.TracerProviderKey).(trace.TracerProvider))
	otel.SetTextMapPropagator(l.ctx.Value(rkmid.PropagatorKey).(propagation.TextMapPropagator))

	//var span trace.Span
	//l.ctx, span = l.ctx.Value(rkmid.TracerKey).(trace.Tracer).Start(l.ctx, "Say")
	//defer span.End()

	metrics.Hello()
	var rsp, err = l.svcCtx.HelloRPC.Say(l.ctx, &hello.SayReq{Name: req.Name})
	if err != nil {
		metrics.HelloFailed()
		return nil, err
	}

	return &types.Rsp{Data: rsp.GetReply(), Result: true, Code: http.StatusOK}, nil
}
