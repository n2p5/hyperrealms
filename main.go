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

// TODO:
// - hyperreals is about creating a simple way to use htmx
// - with smart contracts on gno.land
// - func HyperRender(method, path, headers, data string) string
// - the idea is that a rest call can be translated into a hyper realm call
// - this doesn't have to be used with HTMX, but it could support partial
// - fragment responses.

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

// Ideas
// imagine you are given a subdomain or a path that was "yours"
// storage would be on gnoland directly or using a blessed CDN
// we would have a limited amount of javascript embedded, specifically htmx
// things like webrings would be first-class citizens
// maybe some additional javascript functionality would be allowed, but limited
