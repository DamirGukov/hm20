package models

import (
	"strconv"
)

type Note struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type NoteModel struct {
	notes      []Note
	lastNoteID int
}

func NewNoteModel() *NoteModel {
	return &NoteModel{}
}

func (nm *NoteModel) CreateNote(title, body string) Note {
	nm.lastNoteID++
	id := "note" + strconv.Itoa(nm.lastNoteID)
	note := Note{
		ID:    id,
		Title: title,
		Body:  body,
	}
	nm.notes = append(nm.notes, note)
	return note
}

func (nm *NoteModel) GetAllNotes() []Note {
	return nm.notes
}

func (nm *NoteModel) GetNoteByID(id string) (*Note, bool) {
	for _, note := range nm.notes {
		if note.ID == id {
			return &note, true
		}
	}
	return nil, false
}

func (nm *NoteModel) DeleteNoteByID(id string) bool {
	for i, note := range nm.notes {
		if note.ID == id {
			nm.notes = append(nm.notes[:i], nm.notes[i+1:]...)
			return true
		}
	}
	return false
}
