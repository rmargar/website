package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"github.com/rmargar/website/pkg/application"
	"github.com/rmargar/website/pkg/config"
	"github.com/rmargar/website/pkg/web/representations"
	"github.com/rmargar/website/pkg/web/resources"
)

type Posts struct {
	PostService application.PostService
}

func SetupPosts(r chi.Router, cfg *config.Config, postService application.PostService) {
	tokenAuth := jwtauth.New("HS256", []byte(cfg.JwtSecret), nil)
	postController := Posts{PostService: postService}
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/api/posts", postController.AddPost)
	})
}

func (p Posts) AddPost(w http.ResponseWriter, r *http.Request) {
	var payload resources.PostPayloadJSONApi

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		representations.WriteBadRequestWithErr(w, err)
		return
	}

	if ok, errors := representations.ValidateInputs(payload); !ok {
		representations.WriteValidationResponse([]representations.ClientError{errors}, w)
		return
	}

	createdPost, err := p.PostService.Create(payload.Title, payload.Content, payload.Tags)

	if err != nil {
		representations.WriteBadRequestWithErr(w, err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := resources.BuildCreatedResponse(createdPost)
	json.NewEncoder(w).Encode(response)

}
