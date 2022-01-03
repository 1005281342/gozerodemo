package logic

import (
	"context"
	"github.com/1005281342/gozerodemo/rpc/tasks/task"
	"github.com/hibiken/asynq"

	"github.com/1005281342/gozerodemo/rpc/tasks/internal/svc"
	"github.com/1005281342/gozerodemo/rpc/tasks/tasks"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *tasks.AddReq) (*tasks.AddRsp, error) {
	var (
		cli = task.GetClient()
		a   *asynq.Task
		err error
	)
	if a, err = task.NewPushTask(in.Payload); err != nil {
		return nil, err
	}

	var info *asynq.TaskInfo
	if info, err = cli.Enqueue(a); err != nil {
		return nil, err
	}
	return &tasks.AddRsp{Uuid: info.ID}, nil
}
