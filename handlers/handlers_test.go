package handlers

import (
	"Homework20/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateNoteHandler(t *testing.T) {
	noteData := map[string]string{
		"title": "Test Note",
		"body":  "This is a test note.",
	}
	jsonData, _ := json.Marshal(noteData)
	req, err := http.NewRequest("POST", "/api/notes", bytes.NewReader(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(NewNoteHandler(models.NewNoteModel()).CreateNoteHandler)
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", recorder.Code)
	}

}

func TestGetAllNotesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/notes", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(NewNoteHandler(models.NewNoteModel()).GetAllNotesHandler)
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", recorder.Code)
	}

}

func TestGetNoteByIDHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/notes/note1", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(NewNoteHandler(models.NewNoteModel()).GetNoteByIDHandler)
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", recorder.Code)
	}
}

func TestDeleteNoteByIDHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/api/notes/note1", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(NewNoteHandler(models.NewNoteModel()).DeleteNoteByIDHandler)
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNoContent {
		t.Errorf("Expected status code 204, got %d", recorder.Code)
	}
}
