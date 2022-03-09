FROM golang:1.17.8

ENV GO111MODULE=on \
    GOBIN="/opt/run" \
    GOPROXY="https://goproxy.cn,https://mirrors.aliyun.com/goproxy,direct"

WORKDIR /

COPY ./* /opt/build/

WORKDIR /opt/build/

RUN mkdir -p /opt/run/log &&\
    go env && \
    go mod vendor -v &&\
    go install -ldflags "-X 'main.goVersion=$(go version)' -X 'main.gitHash=$(git rev-parse --short HEAD)' -X 'main.buildTime=$(date "+%Y-%m-%d %H:%M:%S")' " &&\
    echo "build success"


CMD ["/opt/run/gosleep"]
