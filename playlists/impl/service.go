package impl

import (
	"github.com/bgildson/ifood_backend_challenge/base"
	"github.com/bgildson/ifood_backend_challenge/utils"
)

type playlistsService struct {
	temperatureRepository base.TemperatureRepository
	playlistsRepository   base.PlaylistsRepository
}

// NewPlaylistsService : factory used to create a PlaylistsService
func NewPlaylistsService(temperatureRepository base.TemperatureRepository,
	playlistsRepository base.PlaylistsRepository) base.PlaylistsService {
	return playlistsService{
		temperatureRepository: temperatureRepository,
		playlistsRepository:   playlistsRepository,
	}
}

func (s playlistsService) GetPlaylist(city string, latitude float64, longitude float64) (base.Playlist, error) {
	temperature, err := s.temperatureRepository.GetTemperature(city, latitude, longitude)
	if err != nil {
		return nil, err
	}

	genre := utils.ParseTemperatureToGenre(temperature)

	playlist, err := s.playlistsRepository.GetByGenre(genre)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}
