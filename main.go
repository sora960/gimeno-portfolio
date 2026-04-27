package main

import (
	"html/template"
	"log"
	"net/http"
)

// Data structure for the template
type Profile struct {
	Title string
}

func main() {
	// Serve static files (CSS, JS, Images)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Main Route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/base.html", "templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := Profile{Title: "Gimeno Portfolio"}
		tmpl.ExecuteTemplate(w, "base", data)
	})

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
