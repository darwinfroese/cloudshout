package main

import (
	"log"
	"net/http"

	"github.com/darwinfroese/cloudshout/handlers"
	"github.com/darwinfroese/gowt/mux"
)

func main() {
	m := mux.NewMux()

	m.AddRoute("/", handlers.RenderIndex)

	log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", m))
}
