package base

// Track : represents a track
type Track struct {
	Name string `json:"name"`
}

// Playlist : represents a playlist
type Playlist []Track
