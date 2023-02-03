package application

import (
	"fmt"
	"testing"
	"time"

	"github.com/rmargar/website/pkg/domain"
	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

type postRepoMock struct {
	mock.Mock
}

var mockNumPosts int = 5

func GetMockPosts() []*domain.Post {
	var posts []*domain.Post
	var post domain.Post
	for i := 1; i <= mockNumPosts; i++ {
		post = domain.Post{
			ID:       i,
			Title:    fmt.Sprintf("Test_%d", i),
			Author:   "test",
			Content:  fmt.Sprintf("Test_%d", i),
			Tags:     []string{"Life", "Universe"},
			Added:    time.Now(),
			Modified: time.Now(),
		}
		posts = append(posts, &post)
	}
	return posts
}

func (p *postRepoMock) New(post domain.Post) (*domain.Post, error) {
	var id int
	if post.ID == 0 {
		id = 1
	} else {
		id = post.ID
	}
	return &domain.Post{
		ID:       id,
		Title:    post.Title,
		Author:   post.Author,
		Content:  post.Content,
		Tags:     post.Tags,
		Added:    time.Now(),
		Modified: time.Now(),
	}, nil
}

func (p *postRepoMock) GetAll() ([]*domain.Post, error) {
	return GetMockPosts(), nil
}

func (p *postRepoMock) SearchByTitle(title string) ([]*domain.Post, error) {
	return GetMockPosts(), nil
}

func TestPostService_Create(t *testing.T) {
	postService := NewPostService(&postRepoMock{})
	newPost := domain.Post{
		Title:   "Test",
		Content: "Test",
		Tags:    []string{"Golang", "Testing"},
	}
	createdPost, err := postService.Create(newPost.Title, newPost.Content, newPost.Tags)
	if err != nil {
		t.Errorf("PostService.Create() threw error (%v), expected (%v)", err, nil)
	}

	assert.Equal(t, createdPost.Title, newPost.Title)
	assert.Equal(t, createdPost.Content, newPost.Content)
	assert.Equal(t, createdPost.Tags[0], newPost.Tags[0])
	assert.Equal(t, len(createdPost.Tags), len(newPost.Tags))
	assert.Equal(t, createdPost.ID >= 1, true)
}

func TestPostService_GetAll(t *testing.T) {
	postService := NewPostService(&postRepoMock{})
	posts, err := postService.GetAll()
	if err != nil {
		t.Errorf("PostService.GetAll() threw error (%v), expected (%v)", err, nil)
	}
	assert.Equal(t, len(posts), mockNumPosts)
}

func TestPostService_GetOneByTitle(t *testing.T) {

	type args struct {
		numPosts int
		title    string
	}

	tests := []struct {
		name    string
		args    args
		want    domain.Post
		wantErr bool
	}{
		{
			name:    "Should return one",
			args:    args{title: "Test", numPosts: 1},
			wantErr: false,
		},
		{
			name:    "Should return one with error",
			args:    args{title: "Test", numPosts: 1},
			wantErr: true,
		},
		{
			name:    "Should return none",
			args:    args{title: "Test", numPosts: 0},
			wantErr: true,
		},
		{
			name:    "Should return one",
			args:    args{title: "Test", numPosts: 2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		mockNumPosts = tt.args.numPosts
		postService := NewPostService(&postRepoMock{})
		got, err := postService.GetOneByTitle(tt.args.title)
		if tt.args.numPosts > 0 {
			tt.want = *GetMockPosts()[0]
			assert.Equal(t, got.Title, tt.want.Title)
			assert.Equal(t, got.Content, tt.want.Content)
			assert.Equal(t, got.Tags[0], tt.want.Tags[0])
			assert.Equal(t, len(got.Tags), len(tt.want.Tags))
			assert.Equal(t, got.ID, tt.want.ID)
		}
		if err != nil && !tt.wantErr {
			t.Errorf("PostRepoSql.New() error = %v, wantErr %v", err, nil)
			return
		}
	}
}
