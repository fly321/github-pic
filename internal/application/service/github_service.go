package service

import (
	"github-pic/internal/domain/entity"
	"github-pic/internal/domain/repository"
)

// GitHubService 应用层服务，处理GitHub相关业务逻辑
type GitHubService struct {
	githubRepo repository.GitHubRepository
}

// NewGitHubService 创建新的GitHub服务实例
func NewGitHubService(githubRepo repository.GitHubRepository) *GitHubService {
	return &GitHubService{
		githubRepo: githubRepo,
	}
}

// ListRepositories 获取仓库列表
func (s *GitHubService) ListRepositories() ([]*entity.Repository, error) {
	return s.githubRepo.ListRepositories()
}

// UploadImage 上传图片
func (s *GitHubService) UploadImage(repo string, image *entity.Image, content []byte) error {
	return s.githubRepo.UploadImage(repo, image, content)
}

// GetImageURL 获取图片URL
func (s *GitHubService) GetImageURL(repo, path string, useCDN bool) string {
	return s.githubRepo.GetImageURL(repo, path, useCDN)
}
