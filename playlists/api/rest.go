package api

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bgildson/ifood_backend_challenge/base"
	"github.com/bgildson/ifood_backend_challenge/impl"
	"github.com/bgildson/ifood_backend_challenge/utils"
)

var (
	// ErrQueryParamsInvalids : used when query parameters is not valid
	ErrQueryParamsInvalids = errors.New("to complete the request, inform either 'city' or 'lat' and 'lon' parameters")
	// ErrQueryParamLatInvalid : used when the lat param is invalid
	ErrQueryParamLatInvalid = errors.New("'lat' parameter is invalid")
	// ErrQueryParamLonInvalid : used when the lon param is invalid
	ErrQueryParamLonInvalid = errors.New("'lon' parameter is invalid")
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

// RestPlaylistsHandlerQueryParams : used to normalize query params
type RestPlaylistsHandlerQueryParams struct {
	City      string
	Latitude  float64
	Longitude float64
}

// ConvertValuesToRestPlaylistsHandlerQueryParams : used to convert request query params to a normalized type
func ConvertValuesToRestPlaylistsHandlerQueryParams(params url.Values) (*RestPlaylistsHandlerQueryParams, error) {
	if city := params.Get("city"); city != "" {
		return &RestPlaylistsHandlerQueryParams{
			City:      city,
			Latitude:  0,
			Longitude: 0,
		}, nil
	} else if lat, lon := params.Get("lat"), params.Get("lon"); lat != "" || lon != "" {
		latitude, err := strconv.ParseFloat(lat, 64)
		if err != nil {
			return nil, ErrQueryParamLatInvalid
		}
		longitude, err := strconv.ParseFloat(lon, 64)
		if err != nil {
			return nil, ErrQueryParamLonInvalid
		}
		return &RestPlaylistsHandlerQueryParams{
			City:      "",
			Latitude:  latitude,
			Longitude: longitude,
		}, nil
	}
	return nil, ErrQueryParamsInvalids
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	queryParams, err := ConvertValuesToRestPlaylistsHandlerQueryParams(r.URL.Query())
	if err != nil {
		http.Error(w, utils.PrepareHTTPErrorMessage(err.Error()), http.StatusBadRequest)
		return
	}

	playlist, err := h.PlaylistsService.GetPlaylist(queryParams.City, queryParams.Latitude, queryParams.Longitude)
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
