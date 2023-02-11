package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rmargar/website/pkg/application"
	"github.com/rmargar/website/pkg/config"
	"github.com/rmargar/website/pkg/web/html"
)

type Blog struct {
	html    *html.HTML
	service application.PostService
}

func SetupBlog(r chi.Router, cfg *config.Config, postService application.PostService) {
	html, err := html.NewHTML(cfg.HTML)
	if err != nil {
		panic(err)
	}
	blogController := Blog{service: postService, html: html}
	r.Group(func(r chi.Router) {
		r.Get("/blog", blogController.GetBlogIndex)
		r.Get("/blog/{postUrlPath}", blogController.GetPost)
	})
}

func (b *Blog) GetBlogIndex(w http.ResponseWriter, r *http.Request) {
	t := b.html.Templates["index.tpl"]
	allPosts, _ := b.service.GetAll()
	data := html.RenderData{CurrentURL: r.Host + r.URL.Path, Posts: allPosts}
	t.Execute(w, data)
}

func (b *Blog) GetPost(w http.ResponseWriter, r *http.Request) {
	t := b.html.Templates["post.tpl"]
	urlPath := chi.URLParam(r, "postUrlPath")
	post, _ := b.service.GetByUrlPath(urlPath)
	currentURL := r.Host + r.URL.Path
	data := html.RenderPost(post, currentURL)

	t.Execute(w, data)
}
