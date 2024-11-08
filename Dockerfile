# 使用官方的Golang镜像作为基础镜像
FROM golang:latest

# 设置环境变量，使容器内能够使用代理
ENV GOPROXY=https://goproxy.cn

# 设置工作目录
WORKDIR /app

# 将本地的项目文件复制到容器的工作目录中
COPY . /app

# 下载项目依赖（因为有go.mod和go.sum文件）
RUN go mod download

# 构建Go项目
RUN go build -o myblog2

# 暴露你的应用程序监听的端口（你的代码中是80端口）
EXPOSE 80

# 定义容器启动时执行的命令
CMD ["./myblog2"]