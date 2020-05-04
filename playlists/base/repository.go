package base

// PlaylistsRepository : used to define how to implement a playlists repository
type PlaylistsRepository interface {
	GetByGenre(genre Genre) (Playlist, error)
}

// TemperatureRepository : used to define how to implement temperature repository
type TemperatureRepository interface {
	GetByCity(city string) (float64, error)
	GetByLatitudeLongitude(latitude, longitude float64) (float64, error)
}
