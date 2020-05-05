# ifood_backend_challenge
---

This repository contains a project to satisfy an ifood challenge for backend developer. The challenge has the objetive to create a scalable microservice rest api that serves a playlist based in user parameters.

The api expects to receive either `city` or `lat` and `lon` as query params and use the query params to get a place temperature and return a playlist based on this temperature.

_To see the original challenge requirements, [click here](CHALLENGE.md)._

### Running the microservice

The microservice uses the [Open Weather Map](https://openweathermap.org/api) and [Spotify](https://developer.spotify.com/documentation/web-api/quick-start/) apis and require their credentials be informed.

To take Open Weather Map Credentials, follow the link: [https://openweathermap.org/appid](https://openweathermap.org/appid). The Open Weather Map will provide a `API key` that must be passed as **OPEN_WEATHER_MAP_TOKEN** to the microservice.

To take Spotify Credentials, follow the link: [https://developer.spotify.com/documentation/web-api/quick-start/#set-up-your-account](https://developer.spotify.com/documentation/web-api/quick-start/#set-up-your-account). The Spotify will provide `client_id` and `client_secret` that must be passed as **SPOTIFY_CLIENT_ID** and **SPOTIFY_CLIENT_SECRET** respectively to the microservice.

### Running with Docker

1. Certify that you have a docker installation or follow the steps described [here](https://docs.docker.com/get-docker/)

2. In the terminal run the command bellow exchanging <OPEN_WEATHER_MAP_TOKEN>, <SPOTIFY_CLIENT_ID>, <SPOTIFY_CLIENT_SECRET> with your credentials
```
docker run -p 8080:8080 -e PORT=8080 -e OPEN_WEATHER_MAP_TOKEN=<OPEN_WEATHER_MAP_TOKEN> -e SPOTIFY_CLIENT_ID=<SPOTIFY_CLIENT_ID> -e SPOTIFY_CLIENT_SECRET=<SPOTIFY_CLIENT_SECRET> bgildson/ifood_backend_challenge
```

### Running from source code

1. Your environment must be prepared to run GoLang, how described [here](https://golang.org/doc/install)

2. Clone this repository
```
git clone https://github.com/bgildson/ifood_backend_challenge
```

3. In the terminal, walk to the playlists folder in the cloned repository folder and download project dependencies
```
go mod download
```

4. Copy the file containing the environment variables definition
```
cp .env.example .env
```

5. Set the variables in the .env file with the Open Weather Map and Spotify credentials

6. Run the service
```
go run main.go
```

### Running with kubernetes

1. Your environment must be prepared to run kubernetes, how described [here](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
_A good option to run kubernetes locally is using **minikube**, to prepare your environment to run minikube, follow the steps described [here](https://kubernetes.io/docs/tasks/tools/install-minikube/)_

2. Set your credentials on **k8s/playlists-master.yaml** file

3. In the terminal, run the command
```
kubectl apply -f k8s/playlists-master.yaml
```
