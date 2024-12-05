package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/reftch/go-postgres/db"
	"github.com/reftch/go-postgres/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	// router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	// router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	// router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	// router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
