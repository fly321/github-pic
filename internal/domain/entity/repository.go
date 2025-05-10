package entity

// Repository GitHub仓库实体
type Repository struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	CDNURL      string `json:"cdn_url"` // CDN基础URL，如jsDelivr
}

// NewRepository 创建新的仓库实体
func NewRepository(id int64, name, fullName, description string, private bool) *Repository {
	return &Repository{
		ID:          id,
		Name:        name,
		FullName:    fullName,
		Description: description,
		Private:     private,
	}
}

// SetCDNURL 设置CDN URL
func (r *Repository) SetCDNURL(cdnURL string) {
	r.CDNURL = cdnURL
}
