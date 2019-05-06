#源镜像
FROM golang:alpine

#作者
MAINTAINER Pharbers "pqian@pharbers.com"

#LABEL
LABEL NtmPods.version="1.0.5" maintainer="Alex"

# 安装git
RUN apk add --no-cache git gcc musl-dev

# 下载依赖
RUN git clone https://github.com/go-yaml/yaml /go/src/gopkg.in/yaml.v2 && \
    cd /go/src/gopkg.in/yaml.v2 && git checkout tags/v2.2.2 && \
    git clone https://github.com/go-mgo/mgo /go/src/gopkg.in/mgo.v2 && \
    cd /go/src/gopkg.in/mgo.v2 && git checkout -b v2 && \
    git clone https://github.com/PharbersDeveloper/NtmServiceDeploy.git  /go/src/github.com/PharbersDeveloper/NtmServiceDeploy && \
    cd /go/src/github.com/PharbersDeveloper/NtmServiceDeploy && git checkout -b Alex-0301 && \
    git clone https://github.com/PharbersDeveloper/NtmPods.git /go/src/github.com/PharbersDeveloper/NtmPods && \
    cd /go/src/github.com/PharbersDeveloper/NtmPods && git checkout -b Alex-0301


# 设置工程配置文件的环境变量
ENV NTM_HOME /go/src/github.com/PharbersDeveloper/NtmServiceDeploy/deploy-config
ENV GO111MODULE on

# 构建可执行文件
RUN cd /go/src/github.com/PharbersDeveloper/NtmPods && \
    go build -a && go install

# 暴露端口
EXPOSE 31415

# 设置工作目录
WORKDIR /go/bin

ENTRYPOINT ["NtmPods"]