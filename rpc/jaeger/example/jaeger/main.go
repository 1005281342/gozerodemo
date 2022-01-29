package main

import (
	"context"
	"github.com/1005281342/gozerodemo/rpc/jaeger/jaeger"
	"github.com/tal-tech/go-zero/zrpc"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/polaris"
)

func main() {
	var conn, err = zrpc.NewClientWithTarget("polaris://127.0.0.1:8091/jaeger.rpc?wait=14s")
	if err != nil {
		panic(err)
	}
	var cli = jaeger.NewJaegerClient(conn.Conn())
	_, err = cli.Node1(context.Background(), &jaeger.NodeReq{Cnt: 5, Message: "polaris"})
}
