package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, fileName string) {
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[fileName]
	if !ok {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tempCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tempCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return tempCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return tempCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.tmpl")
			if err != nil {
				return tempCache, err
			}
		}

		tempCache[name] = ts
	}
	return tempCache, nil
}
