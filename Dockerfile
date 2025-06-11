FROM golang:1.22 AS builder

COPY . /src
WORKDIR /src

#执行编译 makefile 编译命令
RUN GOPROXY='https://goproxy.cn,direct' \
    GO111MODULE='on' \
    make build

FROM debian:stable-slim

#RUN apt-get update && apt-get install -y --no-install-recommends \
#		ca-certificates  \
#        netbase \
#        && rm -rf /var/lib/apt/lists/ \
#        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
# 声明挂载点
VOLUME /data/conf

CMD ["./server", "-config", "/data/conf/config.yaml"]
