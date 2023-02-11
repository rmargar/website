package application

import (
	"errors"

	"github.com/rmargar/website/pkg/domain"
	"github.com/rmargar/website/pkg/repository"
)

type PostService interface {
	Create(title string, content string, tags []string, summary string, urlPath string) (*domain.Post, error)
	GetOneByTitle(title string) (*domain.Post, error)
	GetAll() ([]domain.Post, error)
	GetByUrlPath(urlPath string) (domain.Post, error)
}

type Posts struct {
	postRepo repository.PostRepository
}

func (p *Posts) Create(title string, content string, tags []string, summary string, urlPath string) (*domain.Post, error) {
	post := domain.Post{Title: title, Content: content, Tags: tags, URLPath: urlPath, Summary: summary}
	return p.postRepo.New(post)
}

func (p *Posts) GetByUrlPath(urlPath string) (domain.Post, error) {
	post, err := p.postRepo.GetOneByUrlPath(urlPath)
	return *post, err
}

func (p *Posts) GetOneByTitle(title string) (*domain.Post, error) {
	var post *domain.Post
	foundPosts, err := p.postRepo.SearchByTitle(title)
	if err != nil {
		return post, err
	}

	switch {
	case len(foundPosts) > 1:
		return &foundPosts[0], errors.New("More than one post was found")
	case len(foundPosts) == 0:
		return post, errors.New("No posts found")
	default:
		return &foundPosts[0], nil
	}
}

func (p *Posts) GetAll() ([]domain.Post, error) {
	return p.postRepo.GetAll()
}

func NewPostService(p repository.PostRepository) *Posts {
	return &Posts{postRepo: p}
}
