# 构建阶段
FROM golang:1.20-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o github-pic .

# 运行阶段
FROM alpine:latest

# 安装ca-certificates，用于HTTPS请求
RUN apk --no-cache add ca-certificates

# 创建非root用户
RUN adduser -D -g '' appuser

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/github-pic /app/

# 复制配置文件和静态资源
COPY --from=builder /app/config.yaml /app/
COPY --from=builder /app/static/ /app/static/

# 使用非root用户运行
USER appuser

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["/app/github-pic"]