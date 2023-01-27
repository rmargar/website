package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rmargar/website/pkg/config"
	"github.com/rmargar/website/pkg/rest/controllers"
)

func NewRouter(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/static/*", controllers.GetStaticFiles)
	r.Get("/", controllers.GetIndexPage)
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/", http.StatusPermanentRedirect)
	})
	r.Post("/static/", controllers.HandlePostForm(&cfg.SmtpConfig, controllers.ParseForm))
	return r
}
