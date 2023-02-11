package repository

import (
	"fmt"

	"github.com/rmargar/website/pkg/domain"
	"github.com/rmargar/website/pkg/orm"
	"gorm.io/gorm"
)

type PostRepository interface {
	New(post domain.Post) (*domain.Post, error)
	GetAll() ([]domain.Post, error)
	SearchByTitle(string) ([]domain.Post, error)
	GetOneByUrlPath(string) (*domain.Post, error)
	GetByTag(string) ([]domain.Post, error)
}

type PostRepoSql struct {
	Db *gorm.DB
}

func (p *PostRepoSql) New(post domain.Post) (*domain.Post, error) {
	postDB := orm.NewPostDB(post)
	result := p.Db.Create(postDB)
	return orm.NewPost(*postDB), result.Error
}

func (p *PostRepoSql) GetAll() ([]domain.Post, error) {
	var posts []orm.Post
	result := p.Db.Find(&posts)
	if result.Error != nil {
		return orm.NewPosts(posts), result.Error
	}
	return orm.NewPosts(posts), nil
}

func (p *PostRepoSql) SearchByTitle(title string) ([]domain.Post, error) {
	var posts []orm.Post
	result := p.Db.Where("title LIKE ?", "%"+title+"%").Find(&posts)
	if result.Error != nil {
		return orm.NewPosts(posts), result.Error
	}
	return orm.NewPosts(posts), nil
}

func (p *PostRepoSql) GetOneByUrlPath(urlPath string) (*domain.Post, error) {
	var post orm.Post
	result := p.Db.Where("url_path = '" + urlPath + "'").First(&post)
	if result.Error != nil {
		return orm.NewPost(post), result.Error
	}
	return orm.NewPost(post), nil
}

func (p *PostRepoSql) GetByTag(tag string) ([]domain.Post, error) {
	var posts []orm.Post
	result := p.Db.Where(fmt.Sprintf("'%s'=any( tags)", tag)).Find(&posts)
	if result.Error != nil {
		return orm.NewPosts(posts), result.Error
	}
	return orm.NewPosts(posts), nil
}

func NewPostRepository(db *gorm.DB) *PostRepoSql {
	return &PostRepoSql{Db: db}
}
