package api_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bgildson/ifood_backend_challenge/api"
	"github.com/bgildson/ifood_backend_challenge/base"
)

func TestConvertValuesToRestPlaylistsHandlerQueryParams(t *testing.T) {
	// success with city
	values := url.Values{"city": {"city"}}
	result, err := api.ConvertValuesToRestPlaylistsHandlerQueryParams(values)

	require.Nil(t, err)

	require.NotNil(t, result)

	require.Equal(t, values["city"][0], result.City)

	// success with lat,lon
	values = url.Values{"lat": {"0"}, "lon": {"0"}}
	result, err = api.ConvertValuesToRestPlaylistsHandlerQueryParams(values)

	require.Nil(t, err)

	require.NotNil(t, result)

	require.Equal(t, 0.0, result.Latitude)

	require.Equal(t, 0.0, result.Longitude)

	// failure with none param informed
	values = url.Values{}

	result, err = api.ConvertValuesToRestPlaylistsHandlerQueryParams(values)

	require.NotNil(t, err)

	require.Equal(t, api.ErrQueryParamsInvalids, err)

	// failure when inform only lat
	values = url.Values{"lat": {"0"}}

	result, err = api.ConvertValuesToRestPlaylistsHandlerQueryParams(values)

	require.NotNil(t, err)

	require.Equal(t, api.ErrQueryParamLonInvalid, err)

	// failure when inform only lon
	values = url.Values{"lon": {"0"}}

	result, err = api.ConvertValuesToRestPlaylistsHandlerQueryParams(values)

	require.NotNil(t, err)

	require.Equal(t, api.ErrQueryParamLatInvalid, err)
}

type PlaylistsServiceMock struct {
	getPlaylist func(city string, latitude float64, longitude float64) (base.Playlist, error)
}

func (s *PlaylistsServiceMock) GetPlaylist(city string, latitude float64, longitude float64) (base.Playlist, error) {
	return s.getPlaylist(city, latitude, longitude)
}

func TestPlaylistsHandlerGet(t *testing.T) {
	track := base.Track{Name: "track"}
	playlistsService := &PlaylistsServiceMock{
		getPlaylist: func(city string, latitude float64, longitude float64) (base.Playlist, error) {
			return base.Playlist{track}, nil
		},
	}
	handler := api.NewRestPlaylistsHandler(playlistsService)
	var contentSuccess []map[string]interface{}
	var contentFailure map[string]interface{}

	// success with city
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	q := url.Values{}
	q.Set("city", "city")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()

	handler.Get(rr, req)

	require.Equal(t, http.StatusOK, rr.Result().StatusCode)

	data, _ := ioutil.ReadAll(rr.Body)

	err := json.Unmarshal(data, &contentSuccess)

	require.Nil(t, err)

	require.Equal(t, track.Name, contentSuccess[0]["name"])

	// success with lat,lon
	req, _ = http.NewRequest(http.MethodGet, "/", nil)
	q = url.Values{}
	q.Set("lat", "0")
	q.Set("lon", "0")
	req.URL.RawQuery = q.Encode()

	rr = httptest.NewRecorder()

	handler.Get(rr, req)

	require.Equal(t, http.StatusOK, rr.Result().StatusCode)

	data, _ = ioutil.ReadAll(rr.Body)

	err = json.Unmarshal(data, &contentSuccess)

	require.Nil(t, err)

	require.Equal(t, track.Name, contentSuccess[0]["name"])

	// failure without query params
	req, _ = http.NewRequest(http.MethodGet, "/", nil)

	rr = httptest.NewRecorder()

	handler.Get(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)

	data, _ = ioutil.ReadAll(rr.Body)

	err = json.Unmarshal(data, &contentFailure)

	require.Nil(t, err)

	require.Equal(t, api.ErrQueryParamsInvalids.Error(), contentFailure["message"])

	// failure when missing lon
	req, _ = http.NewRequest(http.MethodGet, "/", nil)
	q = url.Values{}
	q.Set("lat", "0")
	req.URL.RawQuery = q.Encode()

	rr = httptest.NewRecorder()

	handler.Get(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)

	data, _ = ioutil.ReadAll(rr.Body)

	err = json.Unmarshal(data, &contentFailure)

	require.Nil(t, err)

	require.Equal(t, api.ErrQueryParamLonInvalid.Error(), contentFailure["message"])

	// failure when missing lat
	req, _ = http.NewRequest(http.MethodGet, "/", nil)
	q = url.Values{}
	q.Set("lon", "0")
	req.URL.RawQuery = q.Encode()

	rr = httptest.NewRecorder()

	handler.Get(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)

	data, _ = ioutil.ReadAll(rr.Body)

	err = json.Unmarshal(data, &contentFailure)

	require.Nil(t, err)

	require.Equal(t, api.ErrQueryParamLatInvalid.Error(), contentFailure["message"])
}
