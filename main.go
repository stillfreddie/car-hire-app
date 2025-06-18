package main

import (
    "fmt"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the Car Hire Web App Officially Designed and Configured by Tech Man!")
}

func about(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "About Tech Man Car Hire")
    fmt.Fprintln(w, "--------------------------------------")
    fmt.Fprintln(w, "We are a tech-driven car rental service based in UK, offering reliable and affordable vehicles for all kinds of journeys.")
    fmt.Fprintln(w, "Whether you're heading to a business meeting or planning a weekend getaway, we’ve got the perfect ride for you.")
    fmt.Fprintln(w, "Powered by innovation. Delivered with class.")
    fmt.Fprintln(w, "")
    fmt.Fprintln(w, "Learn more at: https://techmancarhire.com")
}


func contact(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Contact Tech Man Car Hire")
    fmt.Fprintln(w, "--------------------------------------")
    fmt.Fprintln(w, "Email: techman@gmail.com")
    fmt.Fprintln(w, "Phone: +44 783 000 0000")
    fmt.Fprintln(w, "Website: https://techmancarehire.com")
    fmt.Fprintln(w, "Office Hours: Monday – Friday, 9AM to 5PM")
}

}

func admin(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Admin Dashboard: Bookings will be displayed here soon.")
}

func main() {
    http.HandleFunc("/", home)
    http.HandleFunc("/about", about)
    http.HandleFunc("/contact", contact)
    http.HandleFunc("/admin", admin)

    fmt.Println("Server is running on port 8080...")
    http.ListenAndServe(":8080", nil)
}