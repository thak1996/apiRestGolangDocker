package store

import (
    "fmt"
    "myapigo/models"
)

var Items = make(map[string]models.Item)

func GetAllItems() []models.Item {
    items := []models.Item{}
    for _, item := range Items {
        items = append(items, item)
    }
    return items
}

func GetItem(id string) (models.Item, bool) {
    item, found := Items[id]
    return item, found
}

func CreateItem(item models.Item) {
    Items[item.ID] = item
}

func UpdateItem(id string, newItem models.Item) error {
    if _, found := Items[id]; found {
        Items[id] = newItem
        return nil
    }
    return fmt.Errorf("item not found")
}

func DeleteItem(id string) {
    delete(Items, id)
}
