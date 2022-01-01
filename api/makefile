updateApi:
	goctl api go -api demo.api -dir . -style gozero


run:
	go run main.go -f ./etc/demo-api.yaml

swagger:
    goctl api plugin -plugin goctl-swagger="swagger -filename demo.json" -api demo.api -dir .