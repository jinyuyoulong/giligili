FROM golang as build

ENV GOPROXY https://mirrors.aliyun.com/goproxy/

ADD . /giligili

# 设置默认路径，在哪里运行
WORKDIR /giligili

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server




#---分成两部分看
# 二段构建 好处是没有 golang 环境也可以打包
FROM alpine:3.7

ENV MYSQL_DSN=""
ENV REDIS_ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV GIN_MODE="release"
ENV PORT=3000

RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf

WORKDIR /www

# 二段构建 两个镜像之间是有关系的，从build镜像中将文件 copy 到 当前镜像中
COPY --from=build /giligili/api_server /usr/bin/api_server
ADD ./conf /www/conf

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]