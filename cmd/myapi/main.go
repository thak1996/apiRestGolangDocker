package main

import (
    "log"
    "net/http"
    "myapigo/handlers"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/items", handlers.ItemsHandler).Methods("GET", "POST")
    r.HandleFunc("/items/{id}", handlers.ItemHandler).Methods("GET", "PUT", "DELETE")

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
