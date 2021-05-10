####
FROM golang:1.16-alpine AS builder
LABEL stage=datetime_builder

ENV TZ="Asia/Shanghai"

RUN go env -w GOPROXY="https://goproxy.cn,direct"
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache update && apk --no-cache upgrade

WORKDIR /app
COPY . /app
RUN go build -o main main.go

####
FROM alpine

ENV TZ="Asia/Shanghai"

WORKDIR /app
COPY --from=builder /app/main /app/main

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache update && apk --no-cache upgrade && apk --no-cache add tzdata

EXPOSE 8080
CMD ["./main"]
