FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./etc /app/etc

# 在第一阶段添加字体文件
COPY scripts/font/HarmonyOS_Sans_SC_Bold.ttf /app/font/HarmonyOS_Sans_SC_Bold.ttf

RUN go build -ldflags="-s -w" -o /app/captchaimageserver captchaimageserver.go


FROM alpine:latest

# 安装 bash
RUN apk add --no-cache bash

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/captchaimageserver /app/captchaimageserver
COPY --from=builder /app/etc /app/etc
COPY --from=builder /app/font/ /app/font/

ENV CAPTCHA_IMAGE_FONT_PATH=/app/font/HarmonyOS_Sans_SC_Bold.ttf

CMD ["./captchaimageserver", "-f", "etc/captchaimageserver-api.yaml"]
