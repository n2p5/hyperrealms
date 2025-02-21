package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	defaultPort = "9001"
)

func main() {
	port := os.Getenv("HYPERREALMS_PORT")
	if port == "" {
		port = defaultPort
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello again, %s!", r.URL.Path[1:])
	})

	http.ListenAndServe(":"+port, r)
}
