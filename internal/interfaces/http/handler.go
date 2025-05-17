package http

import (
	"fmt"
	"github-pic/internal/application/service"
	"github-pic/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"time"
)

// Handler HTTP请求处理器
type Handler struct {
	githubService *service.GitHubService
}

// NewHandler 创建新的处理器实例
func NewHandler(githubService *service.GitHubService) *Handler {
	return &Handler{githubService: githubService}
}

// ListRepositories 获取仓库列表
func (h *Handler) ListRepositories(c *gin.Context) {
	repos, err := h.githubService.ListRepositories()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, repos)
}

// UploadImage 处理文件上传
func (h *Handler) UploadImage(c *gin.Context) {
	// 获取表单参数
	repo := c.PostForm("repo")
	useCDN := c.PostForm("use_cdn") == "true"

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "文件上传失败"})
		return
	}

	// 读取文件内容
	src, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "无法读取文件"})
		return
	}
	defer src.Close()

	// 生成文件路径
	timestamp := time.Now().Format("20060102150405")
	filePath := fmt.Sprintf("files/%s_%s", timestamp, file.Filename)

	// 创建文件实体
	fileEntity := entity.NewImage(
		file.Filename,
		repo,
		filePath,
		file.Size,
	)

	// 读取文件内容
	content := make([]byte, file.Size)
	if _, err := src.Read(content); err != nil {
		c.JSON(500, gin.H{"error": "读取文件内容失败"})
		return
	}

	// 上传到GitHub
	if err := h.githubService.UploadImage(repo, fileEntity, content); err != nil {
		c.JSON(500, gin.H{"error": "上传到GitHub失败"})
		return
	}

	// 获取访问URL
	fileURL := h.githubService.GetImageURL(repo, filePath, useCDN)

	// 获取客户端IP
	clientIP := c.ClientIP()

	c.JSON(200, gin.H{
		"url":       fileURL,
		"name":      file.Filename,
		"size":      file.Size,
		"client_ip": clientIP,
	})
}
