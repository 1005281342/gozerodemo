package main

import (
	"flag"
	"fmt"

	"github.com/1005281342/gozerodemo/rpc/jaeger/internal/config"
	"github.com/1005281342/gozerodemo/rpc/jaeger/internal/server"
	"github.com/1005281342/gozerodemo/rpc/jaeger/internal/svc"
	"github.com/1005281342/gozerodemo/rpc/jaeger/jaeger"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/polaris"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/jaeger.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewJaegerServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		jaeger.RegisterJaegerServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	// 注册服务
	var err = polaris.RegitserService(polaris.NewPolarisConfig(c.ListenOn,
		polaris.WithServiceName(c.Etcd.Key),
		polaris.WithNamespace("default")))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
