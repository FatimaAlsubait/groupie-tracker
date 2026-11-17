package handlers

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/internal/models"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RenderError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	query := strings.TrimSpace(r.URL.Query().Get("q"))

	artists, err := models.FetchArtists()
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Failed to fetch artists for search")
		return
	}

	var filtered []models.Artist
	lowerQuery := strings.ToLower(query)

	if query == "" {
		filtered = artists
	} else {
		for _, a := range artists {
			matched := false

			// Match band name
			if strings.Contains(strings.ToLower(a.Name), lowerQuery) {
				matched = true
			}

			// Match members
			if !matched {
				for _, member := range a.Members {
					if strings.Contains(strings.ToLower(member), lowerQuery) {
						matched = true
						break
					}
				}
			}

			// Match creation year or first album date
			if !matched {
				if strings.Contains(strconv.Itoa(a.CreationDate), lowerQuery) ||
					strings.Contains(strings.ToLower(a.FirstAlbum), lowerQuery) {
					matched = true
				}
			}

			if matched {
				filtered = append(filtered, a)
			}
		}
	}

	data := models.IndexPageData{
		Query:   query,
		Artists: filtered,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Failed to load index template")
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		RenderError(w, http.StatusInternalServerError, "Failed to render search results")
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = buf.WriteTo(w)
}
