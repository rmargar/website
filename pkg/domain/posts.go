package domain

import "time"

type Post struct {
	ID       int
	Added    time.Time
	Modified time.Time
	Author   string
	Tags     []string
	Title    string
	Content  string
	Summary  string
	URLPath  string
}
