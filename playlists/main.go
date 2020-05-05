package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"

	"github.com/bgildson/ifood_backend_challenge/api"
	"github.com/bgildson/ifood_backend_challenge/impl"
)

func apiAddr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	return ":" + port
}

func main() {
	// reload environment variables
	godotenv.Load()

	// setup resources
	openWeatherMapToken := os.Getenv("OPEN_WEATHER_MAP_TOKEN")
	spotifyClientID := os.Getenv("SPOTIFY_CLIENT_ID")
	spotifyClientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	temperatureRepository := impl.NewTemperatureRepository(http.DefaultClient, openWeatherMapToken)
	spotifyRepository := impl.NewSpotifyRepository(http.DefaultClient, spotifyClientID, spotifyClientSecret)
	playlistsRepository := impl.NewSpotifyPlaylistsRepository(http.DefaultClient, spotifyRepository)
	playlistsService := impl.NewPlaylistsService(temperatureRepository, playlistsRepository)
	handler := api.NewRestPlaylistsHandler(playlistsService)

	router := chi.NewRouter()
	router.Use(
		middleware.Recoverer,
		middleware.Logger,
	)
	router.Get("/", handler.Get)

	// serve api
	addr := apiAddr()
	fmt.Println("Listening on " + addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
