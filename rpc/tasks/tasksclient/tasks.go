// Code generated by goctl. DO NOT EDIT!
// Source: tasks.proto

package tasksclient

import (
	"context"

	"github.com/1005281342/gozerodemo/rpc/tasks/tasks"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddReq    = tasks.AddReq
	AddRsp    = tasks.AddRsp
	CancelReq = tasks.CancelReq
	CancelRsp = tasks.CancelRsp

	Tasks interface {
		Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddRsp, error)
		Cancel(ctx context.Context, in *CancelReq, opts ...grpc.CallOption) (*CancelRsp, error)
	}

	defaultTasks struct {
		cli zrpc.Client
	}
)

func NewTasks(cli zrpc.Client) Tasks {
	return &defaultTasks{
		cli: cli,
	}
}

func (m *defaultTasks) Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddRsp, error) {
	client := tasks.NewTasksClient(m.cli.Conn())
	return client.Add(ctx, in, opts...)
}

func (m *defaultTasks) Cancel(ctx context.Context, in *CancelReq, opts ...grpc.CallOption) (*CancelRsp, error) {
	client := tasks.NewTasksClient(m.cli.Conn())
	return client.Cancel(ctx, in, opts...)
}
