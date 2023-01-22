package main

import (
	"net/http"
	"os"

	"github.com/rmargar/website/pkg/rest"
)

func main() {
	router := rest.NewRouter()
	http.ListenAndServe(os.Getenv(("HTTP_PORT")), router)
}
