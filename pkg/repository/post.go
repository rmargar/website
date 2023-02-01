package repository

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

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
	result := p.Db.Find(&[]orm.Post{})
	if result.Error != nil {
		return posts, result.Error
	}
	rows, err := result.Rows()
	if err != nil {

	}
	return ParsePostsFromRows(rows, posts)
}

func (p *PostRepoSql) SearchByTitle(title string) ([]orm.Post, error) {
	var posts []orm.Post
	result := p.Db.Where("title LIKE ?", "%"+title+"%").Find(&posts)
	if result.Error != nil {
		return posts, result.Error
	}
	return posts, nil
}

func ParsePostsFromRows(rows *sql.Rows, posts []orm.Post) ([]orm.Post, error) {

	for rows.Next() {
		var post orm.Post
		err := rows.Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt, &post.Author, &post.Title, &post.Content, &post.Tags)
		if err != nil {
			log.Errorf("Scan error: %v", err)
			return []orm.Post{}, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
