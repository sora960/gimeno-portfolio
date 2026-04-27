package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os" // Required for port lookup
)

//go:embed templates/* static/*
var content embed.FS

type Profile struct {
	Title string
}

func main() {
	// 1. Serve static files from the EMBEDDED filesystem
	// This ensures your CSS/Images work even if the folders aren't on the server
	staticFiles := http.FileServer(http.FS(content))
	http.Handle("/static/", staticFiles)

	// 2. Load templates from the EMBEDDED filesystem
	tmpl := template.Must(template.ParseFS(content, "templates/base.html", "templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Profile{Title: "Gimeno Portfolio"}
		err := tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			log.Printf("Template error: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	// 3. Dynamic Port Handling for Cloud Providers
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback for your local Arch machine
	}

	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
