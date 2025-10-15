package handlers

import (
	"encoding/json"
	"io"
	"markdown-note-taking--app/services"
	"net/http"
)

// CheckGrammar recibe Markdown y devuelve errores gramaticales
func CheckGrammar(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		http.Error(w, "Contenido vacío o inválido", http.StatusBadRequest)
		return
	}

	issues, err := services.CheckGrammar(string(body))
	if err != nil {
		http.Error(w, "Error al verificar gramática", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(issues)
}
