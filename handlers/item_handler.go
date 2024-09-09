package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wearhouse/models"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT item_id, item_name, user_id FROM item_table")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items[] models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ItemID, &item.ItemName, &item.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}
	json.NewEncoder(w).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemId := params["itemId"]

	var item models.Item
	err := db.QueryRow("SELECT item_id, item_name, user_id FROM item_table WHERE item_id = $1", itemId).Scan(&item.ItemID, &item.ItemName, &item.UserID); 
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(item)
}

func GetItemByUserId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]

	rows, err := db.Query("SELECT item_id, item_name, user_id FROM item_table WHERE user_id = $1", userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ItemID, &item.ItemName, &item.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}
	json.NewEncoder(w).Encode(items)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	_, err := db.Exec("INSERT INTO item_table(item_name, user_id) VALUES ($1, $2)", item.ItemName, item.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Item Successfully Created")
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemId := params["itemId"]

	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("UPDATE item_table SET item_name = $1 WHERE item_id = $2", item.ItemName, itemId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Item with Id %s updated successfully", itemId)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemId := params["itemId"]

	_, err := db.Exec("DELETE FROM item_table WHERE item_id = $1", itemId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Item with Id %s deleted successfully", itemId)
}