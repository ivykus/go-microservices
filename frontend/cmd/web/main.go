package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	port = ":8081"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})
	fmt.Println("starting frontend service on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {

	partials := []string{
		// "./cmd/web/templates/base.layout.html",
		"./cmd/web/templates/header.gohtml",
		"./cmd/web/templates/footer.gohtml",
	}
	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}
	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
