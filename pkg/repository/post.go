package repository

import (
	"github.com/rmargar/website/pkg/domain"
	"github.com/rmargar/website/pkg/orm"
	"gorm.io/gorm"
)

type PostRepository interface {
	New(record orm.Post) (orm.Post, error)
	GetAll() ([]orm.Post, error)
	SearchByTitle(string) ([]orm.Post, error)
}

type PostRepoSql struct {
	Db *gorm.DB
}

func (p *PostRepoSql) New(post domain.Post) (*domain.Post, error) {
	postDB := orm.NewPostDB(post)
	result := p.Db.Create(postDB)
	return orm.NewPost(*postDB), result.Error
}

func (p *PostRepoSql) GetAll() ([]*domain.Post, error) {
	var posts []orm.Post
	result := p.Db.Find(&posts)
	if result.Error != nil {
		return orm.NewPosts(posts), result.Error
	}
	return orm.NewPosts(posts), nil
}

func (p *PostRepoSql) SearchByTitle(title string) ([]*domain.Post, error) {
	var posts []orm.Post
	result := p.Db.Where("title LIKE ?", "%"+title+"%").Find(&posts)
	if result.Error != nil {
		return orm.NewPosts(posts), result.Error
	}
	return orm.NewPosts(posts), nil
}

func NewPostRepository(db *gorm.DB) *PostRepoSql {
	return &PostRepoSql{Db: db}
}
