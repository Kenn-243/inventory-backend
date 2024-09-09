package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/wearhouse/data"
	"github.com/wearhouse/handlers"
)

func main() {
	handlers.SetDB(data.DB)
	router := mux.NewRouter()

	router.HandleFunc("/GetAllUsers", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/GetUser", handlers.GetUser).Methods("POST")
	router.HandleFunc("/CreateUser", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/UpdateUser/{userId}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/DeleteUser/{userId}", handlers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/GetAllItems", handlers.GetItems).Methods("GET")
	router.HandleFunc("/GetItem/{itemId}", handlers.GetItem).Methods("GET")
	router.HandleFunc("/GetItemByUserId/{userId}", handlers.GetItemByUserId).Methods("GET")
	router.HandleFunc("/CreateItem", handlers.CreateItem).Methods("POST")
	router.HandleFunc("/UpdateItem/{itemId}", handlers.UpdateItem).Methods("PUT")
	router.HandleFunc("/DeleteItem/{itemId}", handlers.DeleteItem).Methods("DELETE")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Allow your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", corsHandler))
}