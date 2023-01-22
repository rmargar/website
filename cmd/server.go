package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rmargar/website/pkg/rest"
)

func main() {
	port := "8000" // hardcoded for now
	log.Println(fmt.Sprintf("server listening in port %s", port))
	if err := http.ListenAndServe(":8000", rest.NewRouter()); err != nil {
		fmt.Println("Http server error")
	}
}
