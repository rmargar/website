package repository

import (
	"github.com/rmargar/website/pkg/orm"
	"gorm.io/gorm"
)

type PostRepository interface {
	New(record orm.Post) (orm.Post, error)
}

// UserRepo should i rename it?
type PostRepoSql struct {
	db *gorm.DB
}

func (gorm *PostRepoSql) New(record orm.Post) (orm.Post, error) {
	result := gorm.db.Create(&record)
	return record, result.Error
}
