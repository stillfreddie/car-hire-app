package main

import (
    "html/template"
    "log"
    "net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
    t, err := template.ParseFiles("templates/layout.html", "templates/" + tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := t.Execute(w, nil); err != nil {
        log.Println("Template execution error:", err)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "home")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "about")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "contact")
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "admin")
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/contact", contactHandler)
    http.HandleFunc("/admin", adminHandler)

    // Static files like images and CSS
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    log.Println("🚀 Server running at http://localhost:8080 ...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}