package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "myapigo/handlers"
    _ "myapigo/docs"
    httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
    r := mux.NewRouter()

    r.Use(handlers.LoggingMiddleware)

    r.HandleFunc("/items", handlers.ItemsHandler).Methods("GET", "POST")
    r.HandleFunc("/items/{id}", handlers.ItemHandler).Methods("GET", "PUT", "DELETE")
    
    r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    log.Println("Iniciando o servidor na porta :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
