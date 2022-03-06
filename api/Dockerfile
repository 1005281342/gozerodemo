FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD third_party third_party
ADD go.mod .
ADD go.sum .
RUN go mod download

ADD api/boot.yaml /app/boot.yaml
COPY . .
COPY api/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/demo api/demo.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/boot.yaml /app/boot.yaml
COPY --from=builder /app/demo /app/demo
COPY --from=builder /app/etc /app/etc

CMD ["./demo", "-f", "etc/demo-api.yaml"]
