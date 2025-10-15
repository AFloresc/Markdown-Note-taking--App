package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type GrammarIssue struct {
	Message      string   `json:"message"`
	Offset       int      `json:"offset"`
	Length       int      `json:"length"`
	Replacements []string `json:"replacements"`
}

type grammarResponse struct {
	Matches []GrammarIssue `json:"matches"`
}

func CheckGrammar(text string) ([]GrammarIssue, error) {
	reqBody := map[string]string{
		"text":     text,
		"language": "es",
	}
	jsonBody, _ := json.Marshal(reqBody)

	resp, err := http.Post("https://api.languagetool.org/v2/check", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("respuesta inválida del servidor de gramática")
	}

	var result grammarResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Matches, nil
}
