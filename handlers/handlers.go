package handlers

import (
	"Homework20/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type NoteHandler struct {
	model *models.NoteModel
}

func NewNoteHandler(model *models.NoteModel) *NoteHandler {
	return &NoteHandler{
		model: model,
	}
}

func (nh *NoteHandler) CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var noteData models.Note
	err := json.NewDecoder(r.Body).Decode(&noteData)
	if err != nil {
		respondWithJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	note := nh.model.CreateNote(noteData.Title, noteData.Body)

	respondWithJSON(w, note, http.StatusCreated)
}

func (nh *NoteHandler) GetAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	notesList := nh.model.GetAllNotes()
	respondWithJSON(w, notesList, http.StatusOK)
}

func (nh *NoteHandler) GetNoteByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	note, exists := nh.model.GetNoteByID(params["id"])
	if !exists {
		http.NotFound(w, r)
		return
	}
	respondWithJSON(w, note, http.StatusOK)
}

func (nh *NoteHandler) DeleteNoteByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	deleted := nh.model.DeleteNoteByID(params["id"])
	if !deleted {
		http.NotFound(w, r)
		return
	}
	respondWithJSON(w, nil, http.StatusNoContent)
}

func respondWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Printf("Something went wrong while writing response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
