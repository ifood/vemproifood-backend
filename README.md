# #VemProiFood! - Backend Challenge

Create a micro-service able to accept RESTful requests receiving as parameter
either city name or lat long coordinates and returns a playlist (only track
names is fine) suggestion according to the current temperature.

## Business rules

* If temperature (celcius) is above 30 degrees, suggest tracks for party
* In case temperature is between 15 and 30 degrees, suggest pop music tracks
* If it's a bit chilly (between 10 and 14 degrees), suggest rock music tracks
* Otherwise, if it's freezing outside, suggests classical music tracks

## Hints

You can make usage of OpenWeatherMaps API (https://openweathermap.org) to fetch
temperature data and Spotify (https://developer.spotify.com) to suggest the
tracks as part of the playlist.

- [Spotify API](https://developer.spotify.com/documentation/web-api/quick-start/) (You can use this Client Id: 08c1a6be652e4fdca07f1815bfd167e4)
- [OpenWeatherMaps API](https://home.openweathermap.org/users/sign_up) (You can use this API Key: b77e07f479efe92156376a8b07640ced)

### Sample cities
http://api.openweathermap.org/data/2.5/weather?q=campinas&appid=b77e07f479efe92156376a8b07640ced
http://api.openweathermap.org/data/2.5/weather?q=salvador&appid=b77e07f479efe92156376a8b07640ced
http://api.openweathermap.org/data/2.5/weather?q=brasilia&appid=b77e07f479efe92156376a8b07640ced
http://api.openweathermap.org/data/2.5/weather?q=fortaleza&appid=b77e07f479efe92156376a8b07640ced
http://api.openweathermap.org/data/2.5/weather?q=manaus&appid=b77e07f479efe92156376a8b07640ced

## Non functional requirements

As this service will be a worldwide success, it must be prepared to be fault
tolerant, responsive and resilient.

Use whatever language, tools and frameworks you feel comfortable to, and
briefly elaborate on your solution, architecture details, choice of patterns
and frameworks.

Also, make it easy to deploy/run your service(s) locally (consider using some
container/vm solution for this). Once done, share your code with us.

## Installation and execution instructions

Requirements: [NodeJs](https://nodejs.org/en/), [Docker](https://docs.docker.com/get-docker/), [Docker Compose](https://docs.docker.com/compose/install/)

First, create a `.env` file with the following variables, and fill then with your credentials:

`SPOTIFY_CLIENT_ID,

SPOTIFY_CLIENT_SECRET,

OPEN_WEATHER_API_KEY`

To install the project, all you have to do is run `install.sh`. After that, the project will automatically
run on port [3333](https://localhost:3333).

To run the project without the installation part, just run `dev.sh`. The project will start on port [3333](https://localhost:3333).

If you don't want to run with Docker, you can always run `npm install` to setup the project and `npm start` to run it.

## Endpoints

This microservice has only one endpoint: `/playlist`.
You can choose between inform a city by his name or inform geocoordinates (latitute and longitude), to make the microservice function properly.

The two ways are displayed below:

City name:
`/playlist?city={city_name}`

Geocoordinates:
`/playlist?lat={latitude}&lon={longitude}`

### Contacts

You can contact me at any moment at:

E-mail: `leo.ferreira@dcx.ufpb.br`

LinkedIn: https://www.linkedin.com/in/leoferreiras/
