const dotenv = require('dotenv');
dotenv.config();
module.exports = {
    spotifyApiKey: process.env.SPOTIFY_CLIENT_ID,
    spotifyClientSecret: process.env.SPOTIFY_CLIENT_SECRET,
    openWeatherApiKey: process.env.OPEN_WEATHER_API_KEY
}
