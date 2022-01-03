package logic

import (
	"context"
	"github.com/1005281342/gozerodemo/rpc/tasks/tasks"
	"net/http"

	"github.com/1005281342/gozerodemo/api/internal/svc"
	"github.com/1005281342/gozerodemo/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLogic {
	return AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req types.AddReq) (*types.Rsp, error) {
	var rsp, err = l.svcCtx.TasksRPC.Add(l.ctx, &tasks.AddReq{Name: tasks.TaskType(tasks.TaskType_value[req.Name]),
		Payload: req.Payload})
	if err != nil {
		return nil, err
	}
	return &types.Rsp{Result: true, Code: http.StatusOK, Data: rsp.Uuid}, nil
}
