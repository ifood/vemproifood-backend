const axios = require('axios')
const tuc = require('temp-units-conv')
const SpotifyWebApi = require('spotify-web-api-node')
const config = require('../config/config')

const spotifyApi = new SpotifyWebApi({
    clientId: config.spotifyApiKey,
    clientSecret: config.spotifyClientSecret,
})

spotifyApi.clientCredentialsGrant().then(
    function(data) {
      console.log('The access token expires in ' + data.body['expires_in'])
      console.log('The access token is ' + data.body['access_token'])

      // Save the access token so that it's used in future calls
      spotifyApi.setAccessToken(data.body['access_token'])
    },
    function(err) {
      console.log('Something went wrong when retrieving an access token', err)
    }
  )

module.exports = {
    async getPlaylist(city, apiKey){
        try{
            const res = await axios.get(`http://api.openweathermap.org/data/2.5/weather?q=${city}&appid=${apiKey}`)

            const kelvinTemp = res.data.main['temp']
            const celciusTemp = tuc.k2c(kelvinTemp)

            const getPlaylistType = selectPlaylistType(celciusTemp)
            const playlist = getMusics(getPlaylistType)

            return playlist
        }catch(error){
            console.log(error)
            return error
        }
    }
}

function selectPlaylistType(temp){
    if(temp > 30){
        return 'party'
    }else if(temp < 30 && temp > 15){
        return 'pop'
    }else if(temp < 15 && temp > 10){
        return 'rock'
    }else{
        return 'classical'
    }
}

async function getMusics(type){
    let playlists;
    await spotifyApi.searchPlaylists(type)
    .then(function(data){
        playlists = data.body.playlists.items
    }, function(err){
        playlists = err
    })

    if(playlists[0]){
        let musicData;
        await spotifyApi.getPlaylist(playlists[0].id)
        .then(function (data) {
            musicData = data.body.tracks.items
        }, function (err) {
            musicData = err
        })

        const tracksInfo = musicData.map(track => {
            return { name: track.track.name, artist: track.track.artists[0].name }
        })

        return tracksInfo
    }

    return 'No playlist found! Try again!'
}
