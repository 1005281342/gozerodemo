package svc

import (
	"github.com/1005281342/gozerodemo/api/internal/config"
	"github.com/1005281342/gozerodemo/rpc/hello/helloclient"
	"github.com/1005281342/gozerodemo/rpc/jaeger/jaegerclient"
	"github.com/1005281342/gozerodemo/rpc/tasks/tasksclient"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	HelloRPC  helloclient.Hello
	TasksRPC  tasksclient.Tasks
	JaegerRPC jaegerclient.Jaeger
}

func NewServiceContext(c config.Config) *ServiceContext {

	var helloCli, err = zrpc.NewClient(c.HelloRPC)
	if err != nil {
		panic(err)
	}

	var tasksCli zrpc.Client
	if tasksCli, err = zrpc.NewClient(c.TasksRPC); err != nil {
		panic(err)
	}

	var jaegerCli zrpc.Client
	if jaegerCli, err = zrpc.NewClient(c.JaegerRPC); err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:    c,
		HelloRPC:  helloclient.NewHello(helloCli),
		TasksRPC:  tasksclient.NewTasks(tasksCli),
		JaegerRPC: jaegerclient.NewJaeger(jaegerCli),
	}
}
