package models

// Artist represents a single band/artist from /api/artists
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"relations"`
}

// Concert holds cleaned location names and their corresponding tour dates
type Concert struct {
	Location string
	Dates    []string
}

// ArtistDetail combines the artist info with their formatted concert list
type ArtistDetail struct {
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Concerts     []Concert
}

// IndexPageData is passed to index.html for rendering catalogues and search results
type IndexPageData struct {
	Query   string
	Artists []Artist
}

// Relation represents the JSON response from /api/relation/<id>
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
