package main 

import (
	"html/template"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		data := map[string]string{
			"Title": "My Simple Go Web App",
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
