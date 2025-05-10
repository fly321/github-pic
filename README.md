# GitHub Pic 项目文档

## 一、项目概述
GitHub Pic 是一个基于 Go 语言开发的项目，其核心功能是管理 GitHub 仓库并实现图片上传功能。用户能够将图片上传至指定的 GitHub 仓库，并获取图片的访问 URL，同时支持使用 jsDelivr CDN 进行加速访问。该项目采用模块化设计，具有良好的可维护性和扩展性。

## 二、项目依赖
### （一）Go 版本
项目使用 Go 1.21 版本进行开发。

### （二）直接依赖
- **`github.com/gin-gonic/gin v1.9.1`**：轻量级的 Go Web 框架，用于构建 HTTP 服务器，处理客户端请求。
- **`github.com/google/go-github/v57 v57.0.0`**：与 GitHub API 进行交互的客户端库，方便实现仓库管理和图片上传等功能。
- **`github.com/spf13/viper v1.18.2`**：强大的配置管理库，支持多种配置文件格式，用于读取和管理项目的配置信息。
- **`golang.org/x/oauth2 v0.15.0`**：用于实现 OAuth2 认证的库，确保项目与 GitHub API 交互时的安全性。

### （三）间接依赖
项目存在多个间接依赖，这些依赖由 Go 模块系统自动管理，通常不需要手动干预。部分间接依赖包括：
- `github.com/bytedance/sonic v1.9.1`
- `github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311`
- `github.com/fsnotify/fsnotify v1.7.0`
- 其他...

## 三、项目结构
```
github-pic/
├── internal/
│   ├── domain/
│   │   └── entity/
│   │       ├── repository.go
│   │       └── image.go
│   ├── infrastructure/
│   │   ├── config/
│   │   │   └── config.go
│   │   └── github/
│   │       └── client.go
│   ├── application/
│   │   └── service/
│   │       └── github_service.go
│   └── interfaces/
│       └── http/
│           ├── server.go
│           └── handler.go
├── static/
│   └── index.html
├── go.mod
├── go.sum
└── main.go
```

### 各目录及文件说明
1. **`internal` 目录**：包含项目的核心业务逻辑和基础设施代码。
    - **`domain` 目录**：存放领域实体，如 `repository.go` 定义了 GitHub 仓库实体，`image.go` 定义了图片实体。
    - **`infrastructure` 目录**：包含基础设施层代码，如 `config` 目录下的 `config.go` 负责配置管理，`github` 目录下的 `client.go` 负责与 GitHub API 交互。
    - **`application` 目录**：包含应用服务层代码，如 `github_service.go` 实现了与 GitHub 相关的业务逻辑。
    - **`interfaces` 目录**：包含接口层代码，如 `http` 目录下的 `server.go` 负责启动 HTTP 服务器，`handler.go` 负责处理 HTTP 请求。
2. **`static` 目录**：存放静态文件，如 `index.html` 是项目的前端页面。
3. **`go.mod` 和 `go.sum`**：Go 模块的依赖管理文件，记录了项目的依赖信息。
4. **`main.go`**：项目的入口文件，负责初始化配置并启动 HTTP 服务器。

## 四、配置说明
### （一）配置结构
项目的配置信息存储在 `internal/infrastructure/config/config.go` 文件中，配置结构如下：
```go
// Config 应用配置结构
type Config struct {
    GitHub struct {
        Token string `mapstructure:"token"`
    }
    Server struct {
        Port string `mapstructure:"port"`
    }
}
```
- **`GitHub.Token`**：用于访问 GitHub API 的令牌，需要在配置文件中设置有效的令牌，以确保项目能够正常与 GitHub 进行交互。
- **`Server.Port`**：HTTP 服务器监听的端口，默认值为 `:8080`，可以根据需要进行修改。

### （二）配置初始化
在 `main.go` 文件中，首先会调用 `config.InitConfig()` 函数来初始化配置：
```go
func main() {
    // 初始化配置
    if err := config.InitConfig(); err != nil {
        log.Fatalf("配置初始化失败: %v", err)
    }

    // 启动HTTP服务器
    server := http.NewServer()
    if err := server.Run(); err != nil {
        log.Fatalf("服务器启动失败: %v", err)
    }
}
```
`InitConfig` 函数会读取配置文件 `config.yaml`，如果配置文件不存在，会创建默认配置文件。

## 五、使用方法
### （一）启动项目
1. 确保已经安装 Go 1.21 或更高版本。
2. 在项目根目录下，执行以下命令下载依赖：
```sh
go mod tidy
```
3. 在 `config.yaml` 文件中配置 GitHub 令牌和服务器端口。
4. 启动项目：
```sh
go run main.go
```

### （二）使用前端页面
启动服务器后，客户端可以通过访问 `static/index.html` 页面来使用图片上传功能。页面使用了 Vue.js、Element UI 和 Tailwind CSS 构建，提供了友好的用户界面。

### （三）仓库选择和图片上传
在 `index.html` 页面中，用户可以选择要上传图片的 GitHub 仓库，并选择是否使用 CDN 加速。点击上传按钮后，选择要上传的图片，系统会将图片上传至指定的 GitHub 仓库，并返回图片的访问 URL。

### （四）上传结果展示
上传成功后，页面会展示图片的相关信息，包括文件名、文件大小、上传 IP、图片尺寸和图片 URL，并提供复制链接的功能。

## 六、贡献指南
如果你想为该项目做出贡献，请遵循以下步骤：
1. Fork 该项目到你的 GitHub 账户。
2. 创建一个新的分支：`git checkout -b feature/your-feature`。
3. 在新分支上进行代码修改和功能开发。
4. 提交你的更改：`git commit -m 'Add some feature'`。
5. 推送到你的分支：`git push origin feature/your-feature`。
6. 在 GitHub 上提交 Pull Request，详细描述你的更改内容和目的。

## 七、许可证
该项目使用 [MIT 许可证](LICENSE)，你可以自由地使用、修改和分发该项目，但需要保留许可证声明。