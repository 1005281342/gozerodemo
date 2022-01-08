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

## 监控

参考[使用prometheus + grafana + pushgateway搭建监控可视化系统](https://juejin.cn/post/6844903848230944776)

配置：以resCode为例 `rk_demo_api_resCode{exported_job="demo",instance="pushgateway:9091",job="prom-stack"}`

![](imgs/img.png)

参考:

[一篇文章带你理解和使用Prometheus的指标](https://frezc.github.io/2019/08/03/prometheus-metrics/)

[容器监控实践—PromQL查询解析](https://segmentfault.com/a/1190000018372390)

按分钟进行统计 `increase(rk_demo_api_resCode{exported_job="demo",instance="pushgateway:9091",job="prom-stack"}[1m])`

## 分布式任务队列

https://github.com/hibiken/asynq

docker run --rm --name asynqmon -d -p 8098:8080 -e REDIS_ADDR=172.17.0.1:6379 hibiken/asynqmon

## 调用链

```shell
# 启动jaeger
docker run -d --name jaeger \
    -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
    -p 5775:5775/udp \
    -p 6831:6831/udp \
    -p 6832:6832/udp \
    -p 5778:5778 \
    -p 16686:16686 \
    -p 14268:14268 \
    -p 14250:14250 \
    -p 9411:9411 \
    jaegertracing/all-in-one:1.23
```
参考：https://rkdev.info/cn/docs/bootstrapper/user-guide/grpc-golang/basic/interceptor-tracing/

![img.png](imgs/jaeger.png)
