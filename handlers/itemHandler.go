package handlers

import (
    "encoding/json"
    "net/http"
    "myapigo/models"
    "myapigo/store"
)

// ItemsHandler handles the creation and retrieval of items
// @Summary Get or Create Items
// @Description Get a list of items or create a new item
// @Tags items
// @Accept  json
// @Produce  json
// @Param item body models.Item true "Item to create"
// @Success 201 {object} models.Item
// @Failure 400 {object} models.ErrorResponse
// @Router /items [post]
func ItemsHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        items := store.GetAllItems()
        json.NewEncoder(w).Encode(items)
    case "POST":
        var item models.Item
        if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(models.ErrorResponse{Code: 400, Message: "Invalid input"})
            return
        }
        store.CreateItem(item)
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(item)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}


// ItemHandler handles individual item operations
// @Summary Get, Update, or Delete an Item
// @Description Get, update, or delete a specific item by its ID
// @Tags items
// @Accept  json
// @Produce  json
// @Param id path string true "Item ID"
// @Success 200 {object} models.Item
// @Failure 404 {object} models.ErrorResponse
// @Failure 405 {object} models.ErrorResponse
// @Router /items/{id} [get]
// @Router /items/{id} [put]
// @Router /items/{id} [delete]
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
        w.WriteHeader(http.StatusNoContent)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}
