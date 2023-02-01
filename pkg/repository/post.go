package repository

import (
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

func (p *PostRepoSql) New(record orm.Post) (orm.Post, error) {
	result := p.Db.Create(&record)
	return record, result.Error
}

func (p *PostRepoSql) GetAll() ([]orm.Post, error) {
	var posts []orm.Post
	result := p.Db.Find(&posts)
	if result.Error != nil {
		return posts, result.Error
	}
	return posts, nil
}

func (p *PostRepoSql) SearchByTitle(title string) ([]orm.Post, error) {
	var posts []orm.Post
	result := p.Db.Where("title LIKE ?", "%"+title+"%").Find(&posts)
	if result.Error != nil {
		return posts, result.Error
	}
	return posts, nil
}

func NewPostRepository(db *gorm.DB) *PostRepoSql {
	return &PostRepoSql{Db: db}
}
