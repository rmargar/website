package main

import (
	"fmt"
	"net/http"

	"github.com/rmargar/website/pkg/application"
	"github.com/rmargar/website/pkg/config"
	"github.com/rmargar/website/pkg/database"
	"github.com/rmargar/website/pkg/logging"
	"github.com/rmargar/website/pkg/repository"
	"github.com/rmargar/website/pkg/rest"
	log "github.com/sirupsen/logrus"
)

var services application.Services
var cfg *config.Config

func init() {
	cfg = config.GetConfig()
	db := database.NewGormDB(&cfg.Database)
	database.MigrateUp(database.GetDB(db), &cfg.Database)
	services.PostService = application.NewPostService(repository.NewPostRepository(db))
}

func main() {
	logging.ConfigureLogger()
	log.Info(fmt.Sprintf("Server listening in port %s", cfg.Port))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), rest.NewRouter(cfg, services)); err != nil {
		log.Println("Http server error")
	}
}
