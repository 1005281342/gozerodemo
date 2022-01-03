updateApi:
	goctl api go -api demo.api -dir . -style gozero

ngrok:
	ngrok http 8888

run:
	go run main.go -f ./etc/demo-api.yaml

swagger:
    goctl api plugin -plugin goctl-swagger="swagger -filename demo.json" -api demo.api -dir .