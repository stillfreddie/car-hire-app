package main

import (
    "html/template"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", handler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    println("🚀 Server is running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}