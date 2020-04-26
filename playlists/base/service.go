package base

// PlaylistsService : used to define how to implement playlists service
type PlaylistsService interface {
	GetPlaylist(city string, latitude float64, longitude float64) (Playlist, error)
}
