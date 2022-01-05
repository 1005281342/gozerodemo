package logic

import (
	"context"
	"github.com/1005281342/gozerodemo/rpc/tasks/task"

	"github.com/1005281342/gozerodemo/rpc/tasks/internal/svc"
	"github.com/1005281342/gozerodemo/rpc/tasks/tasks"

	"github.com/tal-tech/go-zero/core/logx"
)

type CancelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLogic {
	return &CancelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelLogic) Cancel(in *tasks.CancelReq) (*tasks.CancelRsp, error) {
	var err = task.DeleteTask(in.GetName().String(), in.GetUuid())
	if err != nil {
		return nil, err
	}
	return &tasks.CancelRsp{}, nil
}
