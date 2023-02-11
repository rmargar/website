package orm

import (
	"time"

	"github.com/lib/pq"
	"github.com/rmargar/website/pkg/domain"
)

// gorm.Model definition
type Post struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Author    string         `gorm:"default:Ricardo"`
	Tags      pq.StringArray `gorm:"type:text[]"`
	Title     string
	Content   string `gorm:"type:text"`
	URLPath   string `gorm:"column:url_path"`
	Summary   string
}

func NewPostDB(p domain.Post) *Post {
	return &Post{
		ID:      p.ID,
		Author:  p.Author,
		Tags:    p.Tags,
		Title:   p.Title,
		Content: p.Content,
		URLPath: p.URLPath,
		Summary: p.Summary,
	}
}

func NewPost(p Post) *domain.Post {
	return &domain.Post{
		ID:       p.ID,
		Added:    p.CreatedAt,
		Modified: p.UpdatedAt,
		Author:   p.Author,
		Tags:     p.Tags,
		Content:  p.Content,
		Title:    p.Title,
		URLPath:  p.URLPath,
		Summary:  p.Summary,
	}
}

func NewPosts(p []Post) []domain.Post {
	posts := make([]domain.Post, 0, len(p))
	for _, post := range p {
		posts = append(posts, *NewPost(post))
	}
	return posts
}
