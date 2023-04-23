package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, fileName string) {
	templ, _ := template.ParseFiles("./templates/" + fileName)
	if err := templ.Execute(w, nil); err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
