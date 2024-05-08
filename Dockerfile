# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
FROM golang:1.18 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies using go modules.
# Allows for layer caching of modules.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
# -o myapp specifies the output file name; adjust as needed.
RUN go build -o /app/myapp

# 基于 Alpine Linux 的轻量级基础镜像
FROM amazonlinux:2 AS builder-socat

# 安装构建所需的工具
RUN yum install -y wget tar gcc make gzip

# 下载并解压 socat 源码
WORKDIR /tmp
RUN wget http://www.dest-unreach.org/socat/download/socat-1.7.4.4.tar.gz && \
    tar -xzvf socat-1.7.4.4.tar.gz

# 进入 socat 源码目录
WORKDIR /tmp/socat-1.7.4.4

# 编译 socat
RUN ./configure && make && make install

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/
# FROM alpine:latest
FROM amazonlinux:2
# RUN apk --no-cache add ca-certificates
RUN yum install -y net-tools && yum clean all

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/myapp /myapp
# 将构建好的 socat 文件复制到新镜像中
COPY --from=builder-socat /usr/local/bin/socat /usr/local/bin/socat

# Expose port vsock 9090 for the application.
EXPOSE 9090

COPY run.sh /run.sh

# Run the web service on container startup.
COPY kmstool_enclave_cli  /usr/local/bin/
COPY libnsm.so  /usr/local/lib/

CMD ["bash","/run.sh"]

# CMD ["sh", "-c", "/run.sh"]

# CMD ["sh", "-c", "/myapp"]