package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	defaultPort = "8080"
)

func main() {
	port := os.Getenv("HYPERREALMS_PORT")
	if port == "" {
		port = defaultPort
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("called test")
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	http.ListenAndServe(":"+port, nil)
}

func foo() string {
	return "Hello, World!"
}
