package impl_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bgildson/ifood_backend_challenge/base"
	"github.com/bgildson/ifood_backend_challenge/impl"
)

type PlaylistsRepositoryMock struct {
	getByGenre func(genre string) (base.Playlist, error)
}

func (r *PlaylistsRepositoryMock) GetByGenre(genre string) (base.Playlist, error) {
	return r.getByGenre(genre)
}

type TemperatureRepositoryMock struct {
	getTemperature func(city string, latitude float64, longitude float64) (float64, error)
}

func (r *TemperatureRepositoryMock) GetTemperature(city string, latitude float64, longitude float64) (float64, error) {
	return r.getTemperature(city, latitude, longitude)
}

type PlaylistsServiceTestSuite struct {
	suite.Suite
	playlistsRepository   base.PlaylistsRepository
	temperatureRepository base.TemperatureRepository
}

func (s *PlaylistsServiceTestSuite) SetupTest() {
	s.playlistsRepository = &PlaylistsRepositoryMock{
		getByGenre: func(genre string) (base.Playlist, error) {
			return base.Playlist{
				{Name: "track 01"},
			}, nil
		},
	}
	s.temperatureRepository = &TemperatureRepositoryMock{
		getTemperature: func(city string, latitude float64, longitude float64) (float64, error) {
			return 315.15, nil
		},
	}
}

func (s *PlaylistsServiceTestSuite) TestPlaylistServiceGetPlaylist() {
	var temperatureRepository base.TemperatureRepository
	var playlistsRepository base.PlaylistsRepository
	var service base.PlaylistsService

	// success
	temperatureRepository = &TemperatureRepositoryMock{
		getTemperature: func(city string, latitude float64, longitude float64) (float64, error) {
			return 315.15, nil
		},
	}
	playlistsRepository = &PlaylistsRepositoryMock{
		getByGenre: func(genre string) (base.Playlist, error) {
			return base.Playlist{{Name: "track 01"}}, nil
		},
	}

	service = impl.NewPlaylistsService(temperatureRepository, playlistsRepository)

	result, err := service.GetPlaylist("", 0, 0)

	s.Nil(err)

	s.NotNil(result)

	// failure in temperature repository
	temperatureRepository = &TemperatureRepositoryMock{
		getTemperature: func(city string, latitude float64, longitude float64) (float64, error) {
			return 0, errors.New("occur an error")
		},
	}
	playlistsRepository = &PlaylistsRepositoryMock{
		getByGenre: func(genre string) (base.Playlist, error) {
			return nil, nil
		},
	}

	service = impl.NewPlaylistsService(temperatureRepository, playlistsRepository)

	result, err = service.GetPlaylist("", 0, 0)

	s.NotNil(err)

	s.Nil(result)

	// failure in playlists repository
	temperatureRepository = &TemperatureRepositoryMock{
		getTemperature: func(city string, latitude float64, longitude float64) (float64, error) {
			return 0, nil
		},
	}
	playlistsRepository = &PlaylistsRepositoryMock{
		getByGenre: func(genre string) (base.Playlist, error) {
			return nil, errors.New("occur an error")
		},
	}

	service = impl.NewPlaylistsService(temperatureRepository, playlistsRepository)

	result, err = service.GetPlaylist("", 0, 0)

	s.NotNil(err)

	s.Nil(result)
}

func TestPlaylistsServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PlaylistsServiceTestSuite))
}
