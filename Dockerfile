# 使用官方 Go 镜像作为基础镜像
FROM golang:1.24

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目文件到容器中
COPY . .

# 构建项目
RUN go build -o server main.go

# 设置静态文件目录
RUN mkdir -p /app/pages

# 复制静态文件到容器中
COPY pages /app/pages

# 暴露端口
EXPOSE 8080

# 设置启动命令
CMD ["./server"]
