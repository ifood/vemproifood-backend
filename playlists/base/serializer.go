package base

// PlaylistSerializer : used to define how to implement a playlist serializer
type PlaylistSerializer interface {
	Decode([]byte) (Playlist, error)
	Encode(Playlist) ([]byte, error)
}
