package handlers

import (
	"bytes"
	"html/template"
	"net/http"
)

type ErrorData struct {
	Code    int
	Message string
}

// RenderError safely buffers and outputs custom error pages without crashing
func RenderError(w http.ResponseWriter, code int, message string) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, message, code)
		return
	}

	var buf bytes.Buffer
	data := ErrorData{Code: code, Message: message}
	if err := tmpl.Execute(&buf, data); err != nil {
		http.Error(w, message, code)
		return
	}

	w.WriteHeader(code)
	_, _ = buf.WriteTo(w)
}
