package logic

import (
	"context"
	"net/http"

	"github.com/1005281342/gozerodemo/api/internal/svc"
	"github.com/1005281342/gozerodemo/api/internal/types"
	"github.com/1005281342/gozerodemo/rpc/tasks/tasks"
	"github.com/tal-tech/go-zero/core/logx"
)

type CancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) CancelLogic {
	return CancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelLogic) Cancel(req types.CancelReq) (*types.Rsp, error) {
	var _, err = l.svcCtx.TasksRPC.Cancel(l.ctx, &tasks.CancelReq{Name: tasks.TaskType(tasks.TaskType_value[req.Name]),
		Uuid: req.Uuid})
	if err != nil {
		return nil, err
	}
	return &types.Rsp{Result: true, Code: http.StatusOK}, nil
}
