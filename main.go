package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	//Routes
	r.HandleFunc("/grammar-check", handlers.CheckGrammar).Methods("POST")
	r.HandleFunc("/notes", handlers.SaveNote).Methods("POST")
	r.HandleFunc("/notes", handlers.ListNotes).Methods("GET")
	r.HandleFunc("/notes/{id}/rendered", handlers.RenderNote).Methods("GET")

	//Start server
	log.Println("Servidor escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
