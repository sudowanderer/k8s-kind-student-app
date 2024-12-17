# 使用多阶段构建，第一阶段为构建镜像
FROM golang:1.23.2-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 以利用缓存
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用程序
RUN go build -o main .

# 第二阶段为运行镜像，使用更小的基础镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 复制构建好的二进制文件
COPY --from=builder /app/main .

# 暴露端口
EXPOSE 8080

# 启动程序
CMD ["./main"]