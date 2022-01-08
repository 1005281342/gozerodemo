updateRpc:
	goctl rpc proto -src jaeger.proto -dir .

buildlinux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main jaeger.go