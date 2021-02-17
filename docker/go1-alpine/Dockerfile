FROM alpine:3

#### config
ARG GIT_REPO=${GIT_REPO}
ARG PROG=${PROG}
ARG BRANCH=master
ARG PORT=8080
ARG ARGS=""
ARG GO_VERSION=${GO_VERSION}

ENV PORT=${PORT}
ENV ARGS=${ARGS}
ENV TZ="Asia/Shanghai"
ENV GIT_TERMINAL_PROMPT="1"

COPY configs/gitconfig       /root/.gitconfig
COPY configs/git-credentials /root/.git-credentials
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

#### maintains
RUN apk --no-cache update  && \
    apk --no-cache upgrade && \
    apk add git tzdata wget libc6-compat

# https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz
RUN tgz=go${GO_VERSION}.linux-amd64.tar.gz        && \
    wget --quiet https://golang.google.cn/dl/$tgz && \
    tar -xf $tgz -C /opt                          && \
    rm $tgz

ENV PATH=/opt/go/bin:$PATH
RUN go env -w GOPROXY="https://goproxy.cn,direct"

#### install program
RUN git clone ${GIT_REPO}.git /opt/${PROG}

COPY Dockerfile /opt/${PROG}/Dockerfile
WORKDIR /opt/${PROG}

RUN git checkout -f ${BRANCH} && go build -o main ./

# go clean -x
RUN go clean -cache -testcache -modcache && \
    rm -r /opt/go /root/go               && \
    apk del git wget

#### setup service
EXPOSE ${PORT}
CMD ["sh", "-c", "./main -port ${PORT} ${ARGS}"]
