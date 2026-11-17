package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"groupie-tracker/internal/handlers"
)

// recoverMW catches unexpected panics and returns a 500 page so the server never crashes
func recoverMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				handlers.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	// 1. Static file server (pure CSS)
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// 2. Application routes
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/artist", handlers.ArtistHandler)
	mux.HandleFunc("/search", handlers.SearchHandler)

	
	// Crash-proof server configuration with timeouts
	server := &http.Server{
		Addr:         ":8080",
		Handler:      recoverMW(mux),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	fmt.Println("Server running at http://localhost:8080")
	// This line uses the 'server' variable defined right above it:
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to run: %v", err)
	}
}
