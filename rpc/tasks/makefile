updateRpc:
	goctl rpc proto -src tasks.proto -dir .

buildlinux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main tasks.go