module github.com/1005281342/gozerodemo

go 1.16

replace github.com/rookie-ninja/rk-zero v0.0.0 => ./third_party/rk-zero

replace github.com/zeromicro/zero-contrib/zrpc/registry/polaris v0.0.0 => ./third_party/zero-contrib/zrpc/registry/polaris

require (
	github.com/fullstorydev/grpchan v1.1.0
	github.com/hibiken/asynq v0.20.0
	github.com/prometheus/client_golang v1.11.0
	github.com/rookie-ninja/rk-entry v1.0.10
	github.com/rookie-ninja/rk-prom v1.1.4
	github.com/rookie-ninja/rk-zero v0.0.0
	github.com/tal-tech/go-zero v1.2.5
	github.com/zeromicro/zero-contrib/zrpc/registry/polaris v0.0.0
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/go-redis/redis/v8 v8.11.4 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11 // indirect
	google.golang.org/genproto v0.0.0-20220126215142-9970aeb2e350 // indirect
	k8s.io/klog/v2 v2.40.1 // indirect
)
