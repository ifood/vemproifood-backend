package impl

import (
	"encoding/json"

	"github.com/bgildson/ifood_backend_challenge/base"
)

// JSONPlaylistSerializer : used to make serializations with playlists
type JSONPlaylistSerializer struct{}

// Decode : receives a list of bytes and serialize to a Playlist
func (p JSONPlaylistSerializer) Decode(data []byte) (base.Playlist, error) {
	var playlist base.Playlist
	err := json.Unmarshal(data, &playlist)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

// Encode : receives a Playlist and serialize to a list of bytes
func (p JSONPlaylistSerializer) Encode(playlist base.Playlist) ([]byte, error) {
	data, err := json.Marshal(playlist)
	if err != nil {
		return nil, err
	}
	return data, nil
}
