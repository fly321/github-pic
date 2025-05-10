package github

import (
	"context"
	"fmt"
	"github-pic/internal/domain/entity"
	"github-pic/internal/infrastructure/config"
	"strings"

	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

// Client GitHub客户端实现
type Client struct {
	client *github.Client
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
	}
}

// ListRepositories 获取用户的仓库列表
func (c *Client) ListRepositories() ([]*entity.Repository, error) {
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
		// 使用jsDelivr CDN
		return fmt.Sprintf("https://cdn.jsdelivr.net/gh/%s/%s", repo, filePath)
	}

	// 使用GitHub原始URL
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/master/%s", repo, filePath)
}
