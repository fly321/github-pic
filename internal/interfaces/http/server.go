package http

import (
	"github-pic/internal/application/service"
	"github-pic/internal/infrastructure/config"
	"github-pic/internal/infrastructure/github"

	"github.com/gin-gonic/gin"
)

// Server HTTP服务器
type Server struct {
	router *gin.Engine
}

// NewServer 创建新的HTTP服务器实例
func NewServer() *Server {
	router := gin.Default()

	// 创建GitHub客户端
	githubClient := github.NewClient()

	// 创建服务实例
	githubService := service.NewGitHubService(githubClient)

	// 创建处理器
	handler := NewHandler(githubService)

	// 配置路由
	router.GET("/api/repositories", handler.ListRepositories)
	router.POST("/api/upload", handler.UploadImage)

	// 配置静态文件服务
	router.Static("/static", "./static")
	router.StaticFile("/", "./static/index.html")

	return &Server{router: router}
}

// Run 启动HTTP服务器
func (s *Server) Run() error {
	return s.router.Run(config.GetConfig().Server.Port)
}
