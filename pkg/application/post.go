package application

import (
	"errors"

	"github.com/rmargar/website/pkg/domain"
	"github.com/rmargar/website/pkg/repository"
)

type PostsService interface {
	Create(title string, content string, tags []string) *domain.Post
}

type Posts struct {
	postRepo repository.PostRepository
}

func (p *Posts) Create(title string, content string, tags []string) (*domain.Post, error) {
	post := domain.Post{Title: title, Content: content, Tags: tags}
	return p.postRepo.New(post)
}

func (p *Posts) GetOneByTitle(title string) (*domain.Post, error) {
	var post *domain.Post
	foundPosts, err := p.postRepo.SearchByTitle(title)
	if err != nil {
		return post, err
	}

	switch {
	case len(foundPosts) > 1:
		return foundPosts[0], errors.New("More than one post was found")
	case len(foundPosts) == 0:
		return post, errors.New("No posts found")
	default:
		return foundPosts[0], nil
	}
}

func (p *Posts) GetAll() ([]*domain.Post, error) {
	return p.postRepo.GetAll()
}

func NewPostService(p repository.PostRepository) *Posts {
	return &Posts{postRepo: p}
}
