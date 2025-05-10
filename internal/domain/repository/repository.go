package repository

import "github-pic/internal/domain/entity"

// GitHubRepository 定义GitHub仓库操作接口
type GitHubRepository interface {
	// ListRepositories 获取用户的仓库列表
	ListRepositories() ([]*entity.Repository, error)

	// UploadImage 上传图片到指定仓库
	UploadImage(repo string, image *entity.Image, content []byte) error

	// GetImageURL 获取图片的访问URL
	GetImageURL(repo, path string, useCDN bool) string
}
