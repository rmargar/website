package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/rmargar/website/pkg/application"
	"github.com/rmargar/website/pkg/config"
	"github.com/rmargar/website/pkg/web/controllers"
)

func NewRouter(cfg *config.Config, services application.Services) *chi.Mux {
	r := chi.NewRouter()
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		JSON: true,
	})

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Recoverer)

	r.Get("/static/*", controllers.GetStaticFiles)
	r.Get("/", controllers.GetIndexPage)
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/", http.StatusPermanentRedirect)
	})
	r.Post("/static/", controllers.HandlePostForm(cfg, controllers.ParseForm))
	controllers.SetupPosts(r, cfg, services.PostService)
	controllers.SetupBlog(r, cfg, services.PostService)
	return r
}
