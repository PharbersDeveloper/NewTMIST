#源镜像
FROM golang:1.12.4-alpine

#作者
MAINTAINER Pharbers "pqian@pharbers.com"

#LABEL
LABEL NtmPods.version="1.0.5" maintainer="Alex"

# 安装git
RUN apk add --no-cache git gcc musl-dev

# 下载依赖
RUN git clone https://github.com/go-yaml/yaml $GOPATH/src/gopkg.in/yaml.v2 && \
    cd $GOPATH/src/gopkg.in/yaml.v2 && git checkout tags/v2.2.2 && \
    git clone https://github.com/go-mgo/mgo $GOPATH/src/gopkg.in/mgo.v2 && \
    cd $GOPATH/src/gopkg.in/mgo.v2 && git checkout -b v2 && \
    git clone -b Alex-0301 https://github.com/PharbersDeveloper/NtmServiceDeploy.git  $GOPATH/src/github.com/PharbersDeveloper/NtmServiceDeploy && \
    git clone -b Alex-0301 https://github.com/PharbersDeveloper/NtmPods.git $GOPATH/src/github.com/PharbersDeveloper/NtmPods

# 设置工程配置文件的环境变量
ENV NTM_HOME $GOPATH/src/github.com/PharbersDeveloper/NtmServiceDeploy/deploy-config
ENV GO111MODULE on

# 构建可执行文件
RUN cd $GOPATH/src/github.com/PharbersDeveloper/NtmPods && \
    go build && go install

# 暴露端口
EXPOSE 31415

# 设置工作目录
WORKDIR $GOPATH/bin

ENTRYPOINT ["Ntm"]