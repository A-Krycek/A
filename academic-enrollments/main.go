package main

import (
	"academic-enrollments/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/enrollments", handlers.EnrollmentsHandler)
	http.HandleFunc("/enrollments/", handlers.EnrollmentHandler)
	http.ListenAndServe(":8080", nil)
}
