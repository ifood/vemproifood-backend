package base

// Track : represents a track
type Track struct {
	Name string `json:"name"`
}

// Playlist : represents a playlist
type Playlist []Track

// Genre : type definition to represents a playlist genre
type Genre string

// Genre's avaliables
const (
	GenreParty     Genre = "party"
	GenrePop       Genre = "pop"
	GenreRock      Genre = "rock"
	GenreClassical Genre = "classical"
)
