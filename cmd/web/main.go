package main

import (
	"net/http"

	"github.com/bash-clevin/go-website/pkg/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(port, nil)
}
