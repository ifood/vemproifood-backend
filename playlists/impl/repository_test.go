package impl_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bgildson/ifood_backend_challenge/base"
	"github.com/bgildson/ifood_backend_challenge/impl"
	"github.com/bgildson/ifood_backend_challenge/utils"
)

type RepositoryTestSuite struct {
	suite.Suite
	spotifyToken        string
	client              utils.HTTPClient
	clientId            string
	clientSecret        string
	genre               base.Genre
	playlist            map[string]interface{}
	openWeatherMapToken string
	city                string
	lat                 float64
	lon                 float64
	temperatureKelvin   float64
	temperatureCelsius  float64
	openWeatherMapData  map[string]interface{}
}

func (s *RepositoryTestSuite) SetupTest() {
	s.spotifyToken = "token"
	s.clientId = "clientId"
	s.clientSecret = "clientSecret"
	s.genre = base.GenreRock
	s.playlist = map[string]interface{}{
		"tracks": map[string]interface{}{
			"items": []map[string]interface{}{
				{"name": "tracks 01"},
				{"name": "tracks 02"},
			},
		},
	}
	s.openWeatherMapToken = "token"
	s.city = "Natal"
	s.lat = 0.1
	s.lon = 0.2
	s.temperatureKelvin = 301.15
	s.temperatureCelsius = 28
	s.openWeatherMapData = map[string]interface{}{
		"main": map[string]interface{}{
			"temp": s.temperatureKelvin,
		},
	}

	s.client = utils.NewHTTPClientMock(
		func(req *http.Request) (*http.Response, error) {
			if req.Method == http.MethodPost && req.URL.String() == "https://accounts.spotify.com/api/token" {
				if req.Header.Get("Authorization") == "" {
					return nil, errors.New("Authorization header not informed")
				}
				if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
					return nil, errors.New("Content-Type header incorrect")
				}
				data := map[string]interface{}{
					"access_token": s.spotifyToken,
					"token_type":   "Bearer",
				}
				content, _ := json.Marshal(data)
				response := &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
					Body:       ioutil.NopCloser(bytes.NewBuffer(content)),
				}
				return response, nil

			} else if req.Method == http.MethodGet && req.URL.String() == "https://api.spotify.com/v1/search?type=track&q=genre:"+string(s.genre) {
				if req.Header.Get("Authorization") == "" {
					return nil, errors.New("Authorization header not informed")
				}
				content, _ := json.Marshal(s.playlist)
				response := &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
					Body:       ioutil.NopCloser(bytes.NewBuffer(content)),
				}
				return response, nil

			} else if req.Method == http.MethodGet &&
				(req.URL.String() == fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", s.city, s.openWeatherMapToken) ||
					req.URL.String() == fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%s", s.lat, s.lon, s.openWeatherMapToken)) {
				content, _ := json.Marshal(s.openWeatherMapData)
				response := &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
					Body:       ioutil.NopCloser(bytes.NewBuffer(content)),
				}
				return response, nil
			}

			return nil, errors.New(fmt.Sprintf("Unexpected request %s: %s", req.Method, req.URL.String()))
		},
	)
}

func (s *RepositoryTestSuite) TestSpotifyRepositoryGetAccessToken() {
	repo := impl.NewSpotifyRepository(s.client, s.clientId, s.clientSecret)

	result, err := repo.GetAccessToken()

	s.Nil(err)

	s.Equal(s.spotifyToken, result)
}

func (s *RepositoryTestSuite) TestSpotifyPlaylistRepositoryGetByGenre() {
	spotifyRepository := impl.NewSpotifyRepository(s.client, s.clientId, s.clientSecret)
	repo := impl.NewSpotifyPlaylistsRepository(s.client, spotifyRepository)

	result, err := repo.GetByGenre(s.genre)

	s.Nil(err)

	for k, v := range result {
		tracks := s.playlist["tracks"].(map[string]interface{})
		items := tracks["items"].([]map[string]interface{})
		s.Equal(items[k]["name"], v.Name)
	}
}

func (s *RepositoryTestSuite) TestTemperatureRepositoryGetByCity() {
	repo := impl.NewTemperatureRepository(s.client, s.openWeatherMapToken)

	result, err := repo.GetByCity(s.city)

	s.Nil(err)

	s.Equal(s.temperatureCelsius, result)
}

func (s *RepositoryTestSuite) TestTemperatureRepositoryGetByLatitudeLongitude() {
	repo := impl.NewTemperatureRepository(s.client, s.openWeatherMapToken)

	result, err := repo.GetByLatitudeLongitude(s.lat, s.lon)

	s.Nil(err)

	s.Equal(s.temperatureCelsius, result)
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
