# GitHub Pic

> 一个简单易用的GitHub图床工具，支持快速上传图片到GitHub仓库并获取图片链接。

## 功能特性

- 🚀 快速上传：支持拖拽或选择图片文件快速上传
- 📦 GitHub存储：利用GitHub仓库作为图床，永久免费存储
- 🔐 安全可靠：使用GitHub Token认证，确保数据安全
- 🌐 链接生成：自动生成图片的直接访问链接
- 📱 响应式设计：支持各种设备访问使用

## 技术栈

- 后端：Go + Gin框架
- 前端：HTML + JavaScript
- 存储：GitHub API
- 配置：Viper

## 快速开始

### 前置要求

- Go 1.16+
- GitHub账号和Personal Access Token

### 安装步骤

1. 克隆项目
```bash
git clone https://github.com/yourusername/github-pic.git
cd github-pic
```

2. 安装依赖
```bash
go mod download
```

3. 配置GitHub Token

在项目根目录创建`config.yaml`文件：
```yaml
github:
  token: "your-github-token"
server:
  port: ":8080"
```

4. 运行项目
```bash
go run main.go
```

访问 http://localhost:8080 即可使用

## 使用说明

1. 打开Web界面
2. 点击上传区域或拖拽图片到上传区域
3. 等待上传完成
4. 复制生成的图片链接

## API接口

### 获取仓库列表

```
GET /api/repositories
```

### 上传图片

```
POST /api/upload
Content-Type: multipart/form-data
```

## 配置说明

配置文件使用YAML格式，支持以下配置项：

- `github.token`: GitHub Personal Access Token
- `server.port`: 服务器监听端口

## 贡献指南

1. Fork 项目
2. 创建新的分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情