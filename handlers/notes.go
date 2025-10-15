package handlers

import (
	"encoding/json"
	"io"
	"markdown-note-taking--app/services"

	"net/http"

	"github.com/gorilla/mux"
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

// RenderNote conversion from Markdown to HTML and retrieves it
func RenderNote(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	note, err := services.GetNoteByID(id)
	if err != nil {
		http.Error(w, "Nota no encontrada", http.StatusNotFound)
		return
	}

	html := services.RenderMarkdown(note.Content)

	resp := map[string]string{"id": id, "html": html}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
