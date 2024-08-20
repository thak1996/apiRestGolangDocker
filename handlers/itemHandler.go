package handlers

import (
    "encoding/json"
    "net/http"
    "myapigo/models"
    "myapigo/store"
)

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        items := store.GetAllItems()
        json.NewEncoder(w).Encode(items)
    case "POST":
        var item models.Item
        json.NewDecoder(r.Body).Decode(&item)
        store.CreateItem(item)
        w.WriteHeader(http.StatusCreated)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func ItemHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/items/"):]
    switch r.Method {
    case "GET":
        if item, found := store.GetItem(id); found {
            json.NewEncoder(w).Encode(item)
        } else {
            w.WriteHeader(http.StatusNotFound)
        }
    case "PUT":
        var item models.Item
        json.NewDecoder(r.Body).Decode(&item)
        if err := store.UpdateItem(id, item); err != nil {
            w.WriteHeader(http.StatusNotFound)
        }
    case "DELETE":
        store.DeleteItem(id)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}
