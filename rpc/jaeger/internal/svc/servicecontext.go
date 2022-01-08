package svc

import (
	"github.com/1005281342/gozerodemo/rpc/hello/helloclient"
	"github.com/1005281342/gozerodemo/rpc/jaeger/internal/config"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	HelloRPC helloclient.Hello
}

func NewServiceContext(c config.Config) *ServiceContext {
	var helloCli, err = zrpc.NewClient(c.HelloRPC)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:   c,
		HelloRPC: helloclient.NewHello(helloCli),
	}
}
