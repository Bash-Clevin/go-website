package handlers

import (
	"net/http"

	"github.com/bash-clevin/go-website/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
