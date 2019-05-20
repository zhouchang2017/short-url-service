# 若运行环境是 Linux 则需把 alpine 换成 debian
# 使用最新版 alpine 作为基础镜像
FROM alpine:latest

# 在容器的根目录下创建 app 目录
RUN mkdir /app

# 将工作目录切换到 /app 下
WORKDIR /app

# 将微服务的服务端可运行文件拷贝到 /app 下
ADD t.wewee /app/t.wewee
ADD client/public /app/client/public
ADD entrypoint.sh /app/entrypoint.sh
ADD logs /app/logs
RUN chmod 777 /app/entrypoint.sh

# 运行服务端
CMD ["./t.wewee"]
