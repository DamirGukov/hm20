package models

import (
	"testing"
)

func TestNoteModel_CreateNote(t *testing.T) {
	model := NewNoteModel()

	note := model.CreateNote("Test Title", "Test Body")

	if note.ID == "" {
		t.Error("Expected note ID not to be empty")
	}
	if note.Title != "Test Title" {
		t.Errorf("Expected title: %s, got: %s", "Test Title", note.Title)
	}
	if note.Body != "Test Body" {
		t.Errorf("Expected body: %s, got: %s", "Test Body", note.Body)
	}
}

func TestNoteModel_GetAllNotes(t *testing.T) {
	model := NewNoteModel()

	notes := model.GetAllNotes()

	if len(notes) != 0 {
		t.Error("Expected notes list to be empty")
	}
}

func TestNoteModel_GetNoteByID_Positive(t *testing.T) {
	model := NewNoteModel()

	note := model.CreateNote("Test Title", "Test Body")
	retrievedNote, exists := model.GetNoteByID(note.ID)

	if !exists {
		t.Error("Expected note to exist")
	}
	if retrievedNote.ID != note.ID {
		t.Errorf("Expected ID: %s, got: %s", note.ID, retrievedNote.ID)
	}
	if retrievedNote.Title != note.Title {
		t.Errorf("Expected title: %s, got: %s", note.Title, retrievedNote.Title)
	}
	if retrievedNote.Body != note.Body {
		t.Errorf("Expected body: %s, got: %s", note.Body, retrievedNote.Body)
	}
}

func TestNoteModel_GetNoteByID_Negative(t *testing.T) {
	model := NewNoteModel()

	_, exists := model.GetNoteByID("nonexistent")

	if exists {
		t.Error("Expected note not to exist")
	}
}

func TestNoteModel_DeleteNoteByID_Positive(t *testing.T) {
	model := NewNoteModel()

	note := model.CreateNote("Test Title", "Test Body")
	deleted := model.DeleteNoteByID(note.ID)

	if !deleted {
		t.Error("Expected note to be deleted")
	}
}

func TestNoteModel_DeleteNoteByID_Negative(t *testing.T) {
	model := NewNoteModel()

	deleted := model.DeleteNoteByID("nonexistent")

	if deleted {
		t.Error("Expected note not to be deleted")
	}
}
