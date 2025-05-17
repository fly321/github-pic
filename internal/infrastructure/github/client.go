package github

import (
	"context"
	"fmt"
	"github-pic/internal/domain/entity"
	"github-pic/internal/infrastructure/config"
	"strings"
	"sync"
	"time"

	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

// repoCache 仓库缓存结构
type repoCache struct {
	data      []*entity.Repository
	timestamp time.Time
	ttl       time.Duration
}

// Client GitHub客户端实现
type Client struct {
	client     *github.Client
	repoCache  *repoCache
	cacheMutex sync.RWMutex
}

// NewClient 创建新的GitHub客户端
func NewClient() *Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GetConfig().GitHub.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return &Client{
		client: github.NewClient(tc),
		repoCache: &repoCache{
			ttl: 5 * time.Minute, // 设置缓存有效期为5分钟
		},
	}
}

// ListRepositories 获取用户的仓库列表（带缓存）
func (c *Client) ListRepositories() ([]*entity.Repository, error) {
	// 尝试从缓存获取
	c.cacheMutex.RLock()
	if c.repoCache.data != nil && time.Since(c.repoCache.timestamp) < c.repoCache.ttl {
		defer c.cacheMutex.RUnlock()
		return c.repoCache.data, nil
	}
	c.cacheMutex.RUnlock()

	// 缓存无效，从GitHub API获取
	ctx := context.Background()
	repos, _, err := c.client.Repositories.List(ctx, "", nil)
	if err != nil {
		return nil, err
	}

	var result []*entity.Repository
	for _, repo := range repos {
		result = append(result, &entity.Repository{
			ID:          repo.GetID(),
			Name:        repo.GetName(),
			FullName:    repo.GetFullName(),
			Description: repo.GetDescription(),
			Private:     repo.GetPrivate(),
		})
	}

	// 更新缓存
	c.cacheMutex.Lock()
	c.repoCache.data = result
	c.repoCache.timestamp = time.Now()
	c.cacheMutex.Unlock()

	return result, nil
}

// UploadImage 上传图片到指定仓库
func (c *Client) UploadImage(repo string, image *entity.Image, content []byte) error {
	ctx := context.Background()

	// 分割仓库全名为所有者和仓库名
	parts := strings.Split(repo, "/")
	if len(parts) != 2 {
		return fmt.Errorf("invalid repository format: %s", repo)
	}
	owner, repoName := parts[0], parts[1]

	// 创建或更新文件
	opts := &github.RepositoryContentFileOptions{
		Message: github.String(fmt.Sprintf("Upload image: %s", image.Name)),
		Content: content,
	}

	_, _, err := c.client.Repositories.CreateFile(ctx, owner, repoName, image.Path, opts)
	return err
}

// GetImageURL 获取图片的访问URL
func (c *Client) GetImageURL(repo, filePath string, useCDN bool) string {
	if useCDN {
		// 输出cdn地址
		// 使用配置文件中的CDN地址
		return fmt.Sprintf(config.GetConfig().GitHub.Cdnurl, repo, filePath)
	}

	// 使用GitHub原始URL
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/master/%s", repo, filePath)
}
