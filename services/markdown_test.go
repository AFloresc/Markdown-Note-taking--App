package services

import (
	"strings"
	"testing"
)

func TestRenderMarkdown(t *testing.T) {
	md := "# Título\n\nEste es un **texto** en *Markdown*."
	html := RenderMarkdown(md)

	if !strings.Contains(html, "<h1>Título</h1>") {
		t.Errorf("No se encontró el encabezado esperado. HTML: %s", html)
	}
	if !strings.Contains(html, "<strong>texto</strong>") {
		t.Errorf("No se encontró el texto en negrita. HTML: %s", html)
	}
	if !strings.Contains(html, "<em>Markdown</em>") {
		t.Errorf("No se encontró el texto en cursiva. HTML: %s", html)
	}
}
