package impl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/bgildson/ifood_backend_challenge/base"
	"github.com/bgildson/ifood_backend_challenge/utils"
)

var (
	// ErrSpotifyAuthentication : used when occur some problem when trying to authenticate on spotify auth api
	ErrSpotifyAuthentication = errors.New("couldn't authenticate on spotify")
	// ErrSpotifyLoadPlaylist : used when occur some problem when trying to load spotify playlist
	ErrSpotifyLoadPlaylist = errors.New("couldn't load playlist")
	// ErrLoadTemperature : used when occur some problem when trying to load temperature
	ErrLoadTemperature = errors.New("couldn't load temperature")
)

// SpotifyRepository : used to implements a repository to make spotify generic operations
type SpotifyRepository interface {
	GetAccessToken() (string, error)
}

type spotifyRepository struct {
	client       utils.HTTPClient
	clientID     string
	clientSecret string
}

// NewSpotifyRepository : used to create a new instance of SpotifyRepository
func NewSpotifyRepository(client utils.HTTPClient, clientID string, clientSecret string) SpotifyRepository {
	return spotifyRepository{
		client:       client,
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

// GetAccessToken : Authenticate on spotify auth server
func (r spotifyRepository) GetAccessToken() (string, error) {
	token := utils.GenerateBasicAuthToken(r.clientID, r.clientSecret)

	authURL, _ := url.Parse("https://accounts.spotify.com/api/token")
	dataValues := url.Values{}
	dataValues.Set("grant_type", "client_credentials")
	request := &http.Request{
		URL:    authURL,
		Method: http.MethodPost,
		Header: http.Header{
			"Authorization": {"Basic " + token},
			"Content-Type":  {"application/x-www-form-urlencoded"},
		},
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte(dataValues.Encode()))),
	}

	response, err := r.client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", ErrSpotifyAuthentication
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var auth struct {
		AccessToken string `json:"access_token"`
	}
	err = json.Unmarshal(data, &auth)
	if err != nil {
		return "", err
	}

	return auth.AccessToken, nil
}

type playlistsRepository struct {
	client            utils.HTTPClient
	spotifyRepository SpotifyRepository
}

// NewSpotifyPlaylistsRepository : used to create a PlaylistsRepository
func NewSpotifyPlaylistsRepository(client utils.HTTPClient, spotifyRepository SpotifyRepository) base.PlaylistsRepository {
	return playlistsRepository{
		client:            client,
		spotifyRepository: spotifyRepository,
	}
}

func (r playlistsRepository) GetByGenre(genre base.Genre) (base.Playlist, error) {
	token, err := r.spotifyRepository.GetAccessToken()
	if err != nil {
		return nil, err
	}

	apiURL, _ := url.Parse("https://api.spotify.com/v1/search?type=track&q=genre:" + string(genre))
	request := http.Request{
		URL:    apiURL,
		Method: http.MethodGet,
		Header: http.Header{
			"Authorization": {"Bearer " + token},
		},
	}

	response, err := r.client.Do(&request)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, ErrSpotifyLoadPlaylist
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var tracks struct {
		Tracks struct {
			Items []struct {
				Name string `json:"name"`
			} `json:"items"`
		} `json:"tracks"`
	}
	if err := json.Unmarshal(body, &tracks); err != nil {
		return nil, err
	}

	var playlist base.Playlist
	for _, v := range tracks.Tracks.Items {
		playlist = append(playlist, base.Track{Name: v.Name})
	}

	return playlist, nil
}

type temperatureRepository struct {
	client   utils.HTTPClient
	apiToken string
}

// NewTemperatureRepository : used to create a TemperatureRepository
func NewTemperatureRepository(client utils.HTTPClient, apiToken string) base.TemperatureRepository {
	return temperatureRepository{
		client:   client,
		apiToken: apiToken,
	}
}

func (r temperatureRepository) getWithURL(url string) (float64, error) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response, err := r.client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return 0, err
	}

	if response.StatusCode != http.StatusOK {
		return 0, ErrLoadTemperature
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var weather struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}
	if err = json.Unmarshal(body, &weather); err != nil {
		return 0, err
	}

	temperature := utils.ParseKelvinToCelsius(weather.Main.Temp)

	return temperature, nil
}

func (r temperatureRepository) GetByCity(city string) (float64, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, r.apiToken)

	return r.getWithURL(url)
}

func (r temperatureRepository) GetByLatitudeLongitude(latitude, longitude float64) (float64, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%s", latitude, longitude, r.apiToken)

	return r.getWithURL(url)
}
