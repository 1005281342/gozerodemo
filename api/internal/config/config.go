package config

import (
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	HelloRPC  zrpc.RpcClientConf
	TasksRPC  zrpc.RpcClientConf
	Namespace string
}
