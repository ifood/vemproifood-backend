const config = require('../config/config')
const service = require('../services/appService')

const apiKey = config.openWeatherApiKey

module.exports={
    async getByCityName(req, res){
        const { city, lat, lon } = req.query

        let musics;
        if(city){
            musics = await service.getPlaylistByCity(city, apiKey)
        }else{
            musics = await service.getPlaylistByGeocoordinates(lat, lon, apiKey)
        }

        return res.json(musics)
    }
}
