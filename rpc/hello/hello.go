package main

import (
	"flag"
	"fmt"
	"github.com/1005281342/gozerodemo/utils/httproxy"
	"github.com/fullstorydev/grpchan"
	"github.com/fullstorydev/grpchan/httpgrpc"
	"github.com/tal-tech/go-zero/core/logx"
	"net"
	"net/http"

	"github.com/1005281342/gozerodemo/rpc/hello/hello"
	"github.com/1005281342/gozerodemo/rpc/hello/internal/config"
	"github.com/1005281342/gozerodemo/rpc/hello/internal/server"
	"github.com/1005281342/gozerodemo/rpc/hello/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/hello.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewHelloServer(ctx)

	reg := grpchan.HandlerMap{}
	httproxy.RegisterHandler(reg, srv, &hello.ServiceDesc)

	var mux http.ServeMux
	httpgrpc.HandleServices(mux.HandleFunc, "/", reg, nil, nil)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.HttpPort))
	if err != nil {
		panic(err)
	}
	logx.Infof("http port: %s", lis.Addr().String())

	httpServer := http.Server{Handler: &mux}
	go httpServer.Serve(lis)
	defer httpServer.Close()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		hello.RegisterHelloServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
