package services

import (
	"errors"
	"sync"

	"markdown-notes-app/models"

	"github.com/google/uuid"
)

var (
	noteStore = make(map[string]models.Note)
	storeLock sync.RWMutex
)

// SaveMarkdownNote saves one Markdown note and returns theID
func SaveMarkdownNote(content string) (string, error) {
	id := uuid.NewString()

	note := models.Note{
		ID:      id,
		Content: content,
	}

	storeLock.Lock()
	defer storeLock.Unlock()
	noteStore[id] = note

	return id, nil
}

// ListMarkdownNotes retieves all the saved notes
func ListMarkdownNotes() []models.Note {
	storeLock.RLock()
	defer storeLock.RUnlock()

	notes := make([]models.Note, 0, len(noteStore))
	for _, note := range noteStore {
		notes = append(notes, note)
	}
	return notes
}

// GetNoteByID retirves one note by ID
func GetNoteByID(id string) (models.Note, error) {
	storeLock.RLock()
	defer storeLock.RUnlock()

	note, exists := noteStore[id]
	if !exists {
		return models.Note{}, errors.New("note not found")
	}
	return note, nil
}
