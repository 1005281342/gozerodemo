updateApi:
	goctl api go -api demo.api -dir . -style gozero

ngrok:
	ngrok http 8888

buildlinux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build demo.go

run:
	go run main.go -f ./etc/demo-api.yaml

swagger:
    goctl api plugin -plugin goctl-swagger="swagger -filename demo.json" -api demo.api -dir .