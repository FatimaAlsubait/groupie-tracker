package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"
)

// Base API endpoint URL
const baseURL = "https://groupietrackers.01edu.ai/api"

var httpClient = &http.Client{Timeout: 10 * time.Second}

// FetchArtists retrieves the complete band list
func FetchArtists() ([]Artist, error) {
	resp, err := httpClient.Get(baseURL + "/artists")
	if err != nil {
		return nil, fmt.Errorf("failed to reach artists API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("artists API responded with status %d", resp.StatusCode)
	}

	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, fmt.Errorf("failed to decode artists JSON: %w", err)
	}
	return artists, nil
}

// FetchArtistDetail fetches an artist and transforms their relations into structured concerts
func FetchArtistDetail(id int) (*ArtistDetail, error) {
	artists, err := FetchArtists()
	if err != nil {
		return nil, err
	}

	var baseArtist *Artist
	for _, a := range artists {
		if a.ID == id {
			baseArtist = &a
			break
		}
	}
	if baseArtist == nil {
		return nil, nil // Artist not found
	}

	// Fetch concert relations for this artist ID
	resp, err := httpClient.Get(fmt.Sprintf("%s/relation/%d", baseURL, id))
	if err != nil {
		return nil, fmt.Errorf("failed to reach relation API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("relation API responded with status %d", resp.StatusCode)
	}

	var rel Relation
	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		return nil, fmt.Errorf("failed to decode relation JSON: %w", err)
	}

	// Convert map[location][]dates into sorted Concert slice
	var concerts []Concert
	for loc, dates := range rel.DatesLocations {
		cleanedLoc := formatLocation(loc)
		concerts = append(concerts, Concert{
			Location: cleanedLoc,
			Dates:    dates,
		})
	}

	// Sort concerts alphabetically by location
	sort.Slice(concerts, func(i, j int) bool {
		return concerts[i].Location < concerts[j].Location
	})

	return &ArtistDetail{
		ID:           baseArtist.ID,
		Image:        baseArtist.Image,
		Name:         baseArtist.Name,
		Members:      baseArtist.Members,
		CreationDate: baseArtist.CreationDate,
		FirstAlbum:   baseArtist.FirstAlbum,
		Concerts:     concerts,
	}, nil
}

// formatLocation transforms "dunedin-new_zealand" into "Dunedin, New Zealand"
func formatLocation(raw string) string {
	parts := strings.Split(raw, "-")
	for i, part := range parts {
		words := strings.Split(part, "_")
		for j, word := range words {
			upper := strings.ToUpper(word)
			if upper == "USA" || upper == "UK" {
				words[j] = upper
			} else if len(word) > 0 {
				words[j] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
			}
		}
		parts[i] = strings.Join(words, " ")
	}
	return strings.Join(parts, ", ")
}
