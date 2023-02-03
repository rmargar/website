package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/go-chi/jwtauth"
	"github.com/rmargar/website/pkg/config"
	"github.com/rmargar/website/pkg/rest/controllers"
)

func NewRouter(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		JSON: true,
	})

	tokenAuth := jwtauth.New("HS256", []byte(cfg.JwtSecret), nil)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Recoverer)

	r.Get("/static/*", controllers.GetStaticFiles)
	r.Get("/", controllers.GetIndexPage)
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/", http.StatusPermanentRedirect)
	})
	r.Post("/static/", controllers.HandlePostForm(&cfg.SmtpConfig, controllers.ParseForm))
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/api/posts", controllers.AddPost)
	})
	return r
}
