package entity

import "time"

// Image 图片实体
type Image struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Size      int64     `json:"size"`
	Repo      string    `json:"repo"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
}

// NewImage 创建新的图片实体
func NewImage(name, repo, path string, size int64) *Image {
	return &Image{
		Name:      name,
		Repo:      repo,
		Path:      path,
		Size:      size,
		CreatedAt: time.Now(),
	}
}
