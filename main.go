package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, fileName string) {
	templ, _ := template.ParseFiles("./templates/" + fileName)
	if err := templ.Execute(w, nil); err != nil {
		fmt.Println("error parsing template", err)
		return
	}

}

func main() {
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(port, nil)
}
