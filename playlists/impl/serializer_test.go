package impl_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bgildson/ifood_backend_challenge/base"
	"github.com/bgildson/ifood_backend_challenge/impl"
)

func TestJSONPlaylistSerializerDecode(t *testing.T) {
	serializer := impl.JSONPlaylistSerializer{}

	// success
	playlist := []map[string]interface{}{
		{"name": "track 01", "artist": map[string]interface{}{"name": "artist name"}},
		{"name": "track 02", "other_field": false},
	}

	data, _ := json.Marshal(playlist)

	result, err := serializer.Decode(data)

	require.Nil(t, err)

	require.Len(t, result, len(playlist))

	for k, v := range playlist {
		require.Equal(t, v["name"], result[k].Name)
	}

	// failure
	data = []byte("something that is not a json")

	result, err = serializer.Decode(data)

	require.NotNil(t, err)

	require.Nil(t, result)
}

func TestJSONPlaylistSerializerEncode(t *testing.T) {
	serializer := impl.JSONPlaylistSerializer{}

	// success
	playlist := base.Playlist{
		{Name: "track 01"},
		{Name: "track 02"},
	}

	result, err := serializer.Encode(playlist)

	require.Nil(t, err)

	var data []map[string]interface{}
	json.Unmarshal(result, &data)

	require.Len(t, data, len(playlist))

	for k, v := range playlist {
		require.Equal(t, v.Name, data[k]["name"])
	}
}
