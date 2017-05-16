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
	m.AddRoute("/main.css", handlers.RenderCSS)
	m.AddRoute("/main.js", handlers.ServeJS)
	m.AddRoute("/admin", handlers.RenderAdmin)
	m.AddRoute("/api/v1/blog", handlers.CreatePostHandler)

	log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", m))
}
