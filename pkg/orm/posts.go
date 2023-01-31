package orm

import (
	"time"
)

// gorm.Model definition
type Post struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Author    string `gorm:"default:Ricardo"`
	Title     string
	Content   string `gorm:"type:text"`
}
