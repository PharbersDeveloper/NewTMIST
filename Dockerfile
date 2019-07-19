#源镜像
FROM golang:1.12.4-alpine

#作者
MAINTAINER Pharbers "pqian@pharbers.com"

# 安装系统级依赖
RUN apk add --no-cache git gcc musl-dev mercurial bash gcc g++ make pkgconfig openssl-dev

# 设置工程配置文件的环境变量
ENV PKG_CONFIG_PATH /usr/lib/pkgconfig
ENV NTM_HOME $GOPATH/src/github.com/PharbersDeveloper/NtmServiceDeploy/deploy-config
ENV BM_KAFKA_CONF_HOME $GOPATH/src/github.com/PharbersDeveloper/NtmServiceDeploy/deploy-config/resource/kafkaconfig.json
ENV BM_XMPP_CONF_HOME $GOPATH/src/github.com/PharbersDeveloper/NtmServiceDeploy/deploy-config/resource/xmppconfig.json
ENV GO111MODULE on

# 以LABEL行的变动(version的变动)来划分(变动以上)使用cache和(变动以下)不使用cache
LABEL NtmPods.version="1.0.7" maintainer="Alex"

# 下载kafka
RUN git clone https://github.com/edenhill/librdkafka.git $GOPATH/librdkafka

WORKDIR $GOPATH/librdkafka
RUN ./configure --prefix /usr  && \
make && \
make install

# 下载依赖
#RUN git clone https://github.com/go-yaml/yaml $GOPATH/src/gopkg.in/yaml.v2 && \
#    cd $GOPATH/src/gopkg.in/yaml.v2 && git checkout tags/v2.2.2 && \
#    git clone https://github.com/go-mgo/mgo $GOPATH/src/gopkg.in/mgo.v2 && \
#    cd $GOPATH/src/gopkg.in/mgo.v2 && git checkout -b v2 && \
#    git clone -b Alex-0301 https://github.com/PharbersDeveloper/NtmServiceDeploy.git  $GOPATH/src/github.com/PharbersDeveloper/NtmServiceDeploy && \
#    git clone -b Alex-0301 https://github.com/PharbersDeveloper/NtmPods.git $GOPATH/src/github.com/PharbersDeveloper/NtmPods

RUN git clone -b Alex-0301 https://github.com/PharbersDeveloper/NtmServiceDeploy.git  $GOPATH/src/github.com/PharbersDeveloper/NtmServiceDeploy && \
    git clone -b Alex-0301 https://github.com/PharbersDeveloper/NtmPods.git $GOPATH/src/github.com/PharbersDeveloper/NtmPods

# 构建可执行文件
RUN cd $GOPATH/src/github.com/PharbersDeveloper/NtmPods && \
    go build && go install

# 暴露端口
EXPOSE 31415

# 设置工作目录
WORKDIR $GOPATH/bin

ENTRYPOINT ["Ntm"]