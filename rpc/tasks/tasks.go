package main

import (
	"flag"
	"fmt"
	"github.com/1005281342/gozerodemo/rpc/tasks/task"
	"github.com/1005281342/gozerodemo/rpc/tasks/tasks"

	"github.com/1005281342/gozerodemo/rpc/tasks/internal/config"
	"github.com/1005281342/gozerodemo/rpc/tasks/internal/server"
	"github.com/1005281342/gozerodemo/rpc/tasks/internal/svc"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/tasks.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewTasksServer(ctx)

	task.Init()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		tasks.RegisterTasksServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
