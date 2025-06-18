package main

import (
    "html/template"
    "log"
    "net/http"
    "path/filepath"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
    layout := filepath.Join("templates", "layout.html")
    page := filepath.Join("templates", tmpl+".html")

    templates, err := template.ParseFiles(layout, page)
    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        log.Println("Template error:", err)
        return
    }

    templates.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        renderTemplate(w, "home")
    })
    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        renderTemplate(w, "about")
    })
    http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
        renderTemplate(w, "contact")
    })
    http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
        renderTemplate(w, "admin")
    })

    // Serve static files (CSS, images)
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    log.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}