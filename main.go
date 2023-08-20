package main

import (
	"fmt"
	"log"
	"net/http"

	"Homework20/handlers"
	"Homework20/models"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	noteModel := models.NewNoteModel()
	noteHandler := handlers.NewNoteHandler(noteModel)

	r.HandleFunc("/api/notes", noteHandler.CreateNoteHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/notes", noteHandler.GetAllNotesHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/notes/{id}", noteHandler.GetNoteByIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/notes/{id}", noteHandler.DeleteNoteByIDHandler).Methods(http.MethodDelete)

	fmt.Println("Server start on port:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
