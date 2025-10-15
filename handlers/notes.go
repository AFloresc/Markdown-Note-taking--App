package handlers

import (
	"encoding/json"
	"io"
	"markdown-notes-app/services"
	"net/http"
)

// SaveNote receive a Markdown, saves nota y returns ID
func SaveNote(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		http.Error(w, "Contenido vacío o inválido", http.StatusBadRequest)
		return
	}

	id, err := services.SaveMarkdownNote(string(body))
	if err != nil {
		http.Error(w, "Error al guardar la nota", http.StatusInternalServerError)
		return
	}

	resp := map[string]string{"id": id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// ListNotes retrievs a list of saved notes
func ListNotes(w http.ResponseWriter, r *http.Request) {
	notes := services.ListMarkdownNotes()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}
