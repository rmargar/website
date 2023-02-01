package orm

import (
	"time"

	"github.com/lib/pq"
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
}
