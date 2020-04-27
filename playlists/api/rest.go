package api

import (
	"net/http"
	"strconv"

	"github.com/bgildson/ifood_backend_challenge/base"
	"github.com/bgildson/ifood_backend_challenge/impl"
	"github.com/bgildson/ifood_backend_challenge/utils"
)

// PlaylistsHandler : used to define how to connect the services to the api
type PlaylistsHandler interface {
	Get(http.ResponseWriter, *http.Request)
}

type handler struct {
	base.PlaylistsService
}

// NewRestPlaylistsHandler : used create a new playlists handler for a rest api
func NewRestPlaylistsHandler(playlistsService base.PlaylistsService) PlaylistsHandler {
	return &handler{
		PlaylistsService: playlistsService,
	}
}

func parseStringToFloat64(value string) float64 {
	if value == "" {
		return 0
	}
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return result
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	city := params.Get("city")
	latitude := parseStringToFloat64(params.Get("lat"))
	longitude := parseStringToFloat64(params.Get("lon"))

	if city == "" && latitude == 0 && longitude == 0 {
		message := "To complete the request, inform either 'city' or 'lat' and 'lon' parameters!"
		http.Error(w, utils.PrepareHTTPErrorMessage(message), http.StatusBadRequest)
		return
	}

	playlist, err := h.PlaylistsService.GetPlaylist(city, latitude, longitude)
	if err != nil {
		http.Error(w, utils.PrepareHTTPErrorMessage(err.Error()), http.StatusBadRequest)
		return
	}

	data, err := impl.JSONPlaylistSerializer{}.Encode(playlist)
	if err != nil {
		http.Error(w, utils.PrepareHTTPErrorMessage(err.Error()), http.StatusBadRequest)
		return
	}

	w.Write(data)
}
