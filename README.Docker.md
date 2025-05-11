# Docker 使用指南

本文档提供了如何使用 Docker 构建和运行 GitHub-Pic 应用的说明。

## 前提条件

- 安装 [Docker](https://www.docker.com/get-started)
- 安装 [Docker Compose](https://docs.docker.com/compose/install/)（可选，用于使用docker-compose.yml）

## 使用 Docker 构建和运行

### 方法一：直接使用 Docker 命令

1. 构建 Docker 镜像

```bash
docker build -t github-pic .
```

2. 运行 Docker 容器

```bash
docker run -p 8080:8080 -v $(pwd)/config.yaml:/app/config.yaml github-pic
```

### 方法二：使用 Docker Compose

1. 使用 Docker Compose 启动服务

```bash
docker-compose up -d
```

2. 查看日志

```bash
docker-compose logs -f
```

3. 停止服务

```bash
docker-compose down
```

## 配置说明

### GitHub Token 配置

您需要在 `config.yaml` 文件中配置有效的 GitHub Token，或者通过环境变量注入。

如果您想通过环境变量注入 GitHub Token，请修改 docker-compose.yml 文件，取消环境变量部分的注释并设置您的 Token。

## 注意事项

1. 默认情况下，应用将在 8080 端口运行
2. 配置文件通过卷挂载方式提供，您可以在不重新构建镜像的情况下修改配置
3. 请确保您的 GitHub Token 有足够的权限来访问和操作您的仓库

## 自定义构建

如果您需要自定义构建过程，可以修改 Dockerfile。该文件使用多阶段构建来减小最终镜像的大小。