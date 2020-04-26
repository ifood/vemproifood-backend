package base

// PlaylistsRepository : used to define how to implement a playlists repository
type PlaylistsRepository interface {
	GetByGenre(genre string) (Playlist, error)
}

// TemperatureRepository : used to define how to implement temperature repository
type TemperatureRepository interface {
	GetTemperature(city string, latitude float64, longitude float64) (float64, error)
}
