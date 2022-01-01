# README

## RPC
https://go-zero.dev/cn/goctl-rpc.html


## 更新API后生成代码

`goctl api go -api gozerodemo.api -dir . -style gozero`

## 启动服务

`go run gozerodemo.go -f etc/gozerodemo-api.yaml`

## 验证

```
% curl -i -X GET \
    http://localhost:8888/from/me 
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 07 Sep 2021 13:22:35 GMT
Content-Length: 23

{"Message":"hello: me"} 
```