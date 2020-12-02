FROM golang:alpine

ADD configs /root/configs

RUN cp /etc/apk/repositories /etc/apk/repositories.bk && \
  sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk upgrade && apk add git tzdata

# set TZ(timezone) and GOPROXY for golang GOPROXY="https://goproxy.cn,direct"
ENV TZ="Asia/Shanghai" 

RUN go env -w GOPROXY="https://goproxy.cn,direct"

# set timezone for alpine
RUN test -f /etc/localtime && cp /etc/localtime /etc/localtime.bk && \
    test -f /etc/timezone && cp /etc/timezone /etc/timezone.bk && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# add user rover
RUN adduser -D rover

# copy git configs and credential
ENV GIT_TERMINAL_PROMPT=1

RUN cp /root/configs/gitconfig /root/.gitconfig && \
    cp /root/configs/git-credentials /root/.git-credentials

RUN cp /root/configs/gitconfig /home/rover/.gitconfig && \
    cp /root/configs/git-credentials /home/rover/.git-credentials && \
    chown rover:rover /home/rover/.gitconfig /home/rover/.git-credentials

WORKDIR /home/rover
