updateRpc:
	goctl rpc proto -src hello.proto -dir .

buildlinux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main hello.go