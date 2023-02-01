package main

import (
	"fmt"
	"net/http"

	"github.com/rmargar/website/pkg/config"
	"github.com/rmargar/website/pkg/database"
	"github.com/rmargar/website/pkg/logging"
	"github.com/rmargar/website/pkg/rest"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.GetConfig()
	logging.ConfigureLogger()
	log.Info(fmt.Sprintf("Server listening in port %s", cfg.Port))

	db := database.NewDB(&cfg.Database)

	database.MigrateUp(db, &cfg.Database)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), rest.NewRouter(cfg)); err != nil {
		log.Println("Http server error")
	}
}
