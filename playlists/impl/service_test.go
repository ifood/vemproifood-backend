package impl_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bgildson/ifood_backend_challenge/base"
	"github.com/bgildson/ifood_backend_challenge/impl"
)

type PlaylistsRepositoryMock struct {
	getByGenre func(genre base.Genre) (base.Playlist, error)
}

func (r *PlaylistsRepositoryMock) GetByGenre(genre base.Genre) (base.Playlist, error) {
	return r.getByGenre(genre)
}

type TemperatureRepositoryMock struct {
	getByCity              func(city string) (float64, error)
	getByLatitudeLongitude func(latitude float64, longitude float64) (float64, error)
}

func (r *TemperatureRepositoryMock) GetByCity(city string) (float64, error) {
	return r.getByCity(city)
}

func (r *TemperatureRepositoryMock) GetByLatitudeLongitude(latitude, longitude float64) (float64, error) {
	return r.getByLatitudeLongitude(latitude, longitude)
}

func TestPlaylistsService(t *testing.T) {
	temperatureRepositorySuccess := &TemperatureRepositoryMock{
		getByCity: func(city string) (float64, error) {
			return 315.15, nil
		},
		getByLatitudeLongitude: func(latitude, longitude float64) (float64, error) {
			return 310.15, nil
		},
	}
	playlistsRepositorySuccess := &PlaylistsRepositoryMock{
		getByGenre: func(genre base.Genre) (base.Playlist, error) {
			return base.Playlist{{Name: "track 01"}}, nil
		},
	}
	temperatureRepositoryFailure := &TemperatureRepositoryMock{
		getByCity: func(city string) (float64, error) {
			return 0, errors.New("occur an error")
		},
		getByLatitudeLongitude: func(latitude, longitude float64) (float64, error) {
			return 0, errors.New("occur an error")
		},
	}
	playlistsRepositoryFailure := &PlaylistsRepositoryMock{
		getByGenre: func(genre base.Genre) (base.Playlist, error) {
			return nil, errors.New("occur an error")
		},
	}
	var service base.PlaylistsService

	// success
	service = impl.NewPlaylistsService(temperatureRepositorySuccess, playlistsRepositorySuccess)

	// using city
	result, err := service.GetPlaylist("city", 0, 0)

	require.Nil(t, err)

	require.NotNil(t, result)

	// using lat,lon
	result, err = service.GetPlaylist("", 0, 0)

	require.Nil(t, err)

	require.NotNil(t, result)

	// failure in temperature repository
	service = impl.NewPlaylistsService(temperatureRepositoryFailure, playlistsRepositorySuccess)

	// using city
	result, err = service.GetPlaylist("city", 0, 0)

	require.NotNil(t, err)

	require.Nil(t, result)

	// using lat,lon
	result, err = service.GetPlaylist("", 0, 0)

	require.NotNil(t, err)

	require.Nil(t, result)

	// failure in playlists repository
	service = impl.NewPlaylistsService(temperatureRepositorySuccess, playlistsRepositoryFailure)

	result, err = service.GetPlaylist("city", 0, 0)

	require.NotNil(t, err)

	require.Nil(t, result)
}
